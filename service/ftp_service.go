package service

import (
	"Training/go-ftp-postgre/config"
	"Training/go-ftp-postgre/domain"
	"bufio"
	"bytes"
	"crypto/rand"
	"fmt"
	"log"
	"math/big"
	"strconv"
	"strings"
	"sync"
	"time"
)

func UploadFile(data string, host string, port int, user string, pass string, fileName string, filePath string) (bool, error) {
	ftp, err := config.SetupConnectionFTP(user, pass, host, port)

	if err != nil {
		return false, err
	}

	_, err = ftp.NameList(filePath)

	if err != nil {
		// Create Folder if status 550
		if strings.Contains(err.Error(), "550") {
			err = ftp.MakeDir(filePath)

			if err != nil {
				return false, err
			}
		} else {
			return false, err
		}
	}

	convert := bytes.NewBufferString(data)

	err = ftp.Stor(filePath + "/" + fileName, convert)

	if err != nil {
		return false, err
	}

	if err := ftp.Quit(); err != nil {
		return false, err
	}

	log.Printf("Upload file %s successfully\n", fileName)

	return true, nil
}

func ReadFile(host string, port int, user string, pass string, fileName string, filePath string) (*[]domain.ProductDTO, error) {
	var mutex sync.Mutex
	wg := &sync.WaitGroup{}

	ftp, err := config.SetupConnectionFTP(user, pass, host, port)

	if err != nil {
		return nil, err
	}

	data, err := ftp.NameList(filePath)

	if err != nil {
		return nil, err
	}

	var products []domain.ProductDTO
	for _, val := range data {
		if strings.Contains(val, fmt.Sprintf(`%s`, fileName)) {
			go func() {
				defer wg.Done()
				wg.Add(1)
				mutex.Lock()

				r, err := ftp.Retr(fmt.Sprintf(`%s/%s`, filePath, val))

				if err != nil {
					log.Fatal(err)
				}

				defer r.Close()

				scanner := bufio.NewScanner(r)
				scanner.Split(bufio.ScanLines)

				var text []string

				for scanner.Scan() {
					text = append(text, scanner.Text())
				}

				r.Close()

				for i, each_ln := range text {
					if i != 0 {
						var product domain.ProductDTO
						split := strings.Split(each_ln, "\t")

						product.ProductCode = split[0]
						product.ProductName = split[1]
						product.ProductSlug = split[2]
						product.ProductDescription = split[3]
						qty, _ := strconv.ParseUint(split[4], 10, 64)
						product.Qty = uint(qty)
						minQty, _ := strconv.ParseUint(split[5], 10, 64)
						product.MinQty =  uint(minQty)
						maxQty, _ := strconv.ParseUint(split[6], 10, 64)
						product.MaxQty = uint(maxQty)
						weight, _ := strconv.ParseUint(split[7], 10, 64)
						product.Weight = uint(weight)
						volume, _ := strconv.ParseUint(split[8], 10, 64)
						product.Volume = uint(volume)

						products = append(products, product)
					}
				}
				mutex.Unlock()
			}()
			time.Sleep(1 * time.Second)
		}
	}

	wg.Wait()

	if err := ftp.Quit(); err != nil {
		log.Fatal(err)
	}

	return &products, nil
}

func MoveFile(host string, port int, user string, pass string, fileName string, filePath string) (bool, error) {
	ftp, err := config.SetupConnectionFTP(user, pass, host, port)

	if err != nil {
		return false, err
	}

	data, err := ftp.NameList(filePath)

	if err != nil {
		return false, err
	}

	for _, val := range data {
		if strings.Contains(val, fmt.Sprintf("%s", fileName)) {
			fileBefore := "/" + filePath + "/" + val
			fileAfter := "/" + filePath  + "/HISTORY/" + val

			err = ftp.Rename(fileBefore, fileAfter)

			if err != nil {
				// Check if file already exists
				if strings.Contains(err.Error(), "file already exists") {
					checkDot := strings.Split(fileAfter, ".")
					n, _ := rand.Int(rand.Reader, big.NewInt(100))

					newName := fmt.Sprintf("%s_%v.%s", checkDot[0], n, checkDot[1])

					err = ftp.Rename(fileBefore, newName)

					if err != nil {
						return false, err
					}
					log.Println("Moving file " + newName + " successfully")
				} else {
					return false, err
				}
			} else {
				log.Println("Moving file " + fileBefore + " successfully")
			}
		}
	}

	return true, nil
}

func RenameFile(host string, port int, user string, pass string, fileName string, filePath string) (bool, error) {
	ftp, err := config.SetupConnectionFTP(user, pass, host, port)

	if err != nil {
		return false, err
	}

	data, err := ftp.NameList(filePath)

	if err != nil {
		return false, err
	}

	for _, val := range data {
		if strings.Contains(val, fmt.Sprintf("%s", fileName)) {
			fileBefore := "/" + filePath + "/" + val

			checkDot := strings.Split(val, ".")

			newName := fmt.Sprintf("RENAME_%s.%s", checkDot[0], checkDot[1])

			fileAfter := "/" + filePath + "/" + newName

			err = ftp.Rename(fileBefore, fileAfter)

			if err != nil {
				return false, err
			} else {
				log.Println("Rename file to " + fileAfter + " successfully")
			}
		}
	}

	return true, nil
}

func DeleteFile(host string, port int, user string, pass string, fileName string, filePath string) (bool, error) {
	ftp, err := config.SetupConnectionFTP(user, pass, host, port)

	if err != nil {
		return false, err
	}

	data, err := ftp.NameList(filePath)

	if err != nil {
		return false, err
	}

	for _, val := range data {
		if strings.Contains(val, fmt.Sprintf("%s", fileName)) {
			file := "/" + filePath + "/" + val
			err = ftp.Delete(file)

			if err != nil {
				return false, err
			} else {
				log.Println("Delete file " + val + " successfully")
			}
		}
	}

	return true, nil
}