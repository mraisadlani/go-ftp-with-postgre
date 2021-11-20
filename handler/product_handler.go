package handler

import (
	"Training/go-ftp-postgre/config"
	"Training/go-ftp-postgre/domain"
	"Training/go-ftp-postgre/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type ProductHandler struct {
	productRepo *domain.ProductRepositoryImpl
}

func BuildProductHandler(productRepo *domain.ProductRepositoryImpl) *ProductHandler {
	return &ProductHandler{
		productRepo: productRepo,
	}
}

// @Summary Upload File
// @Description Upload file to FTP
// @Accept json
// @Produce json
// @Tags FTP Controller
// @Success 200 {object} domain.Message
// @Failure 500,400,404,403 {object} domain.Message
// @Router /UploadFile [post]
func (h *ProductHandler) UploadFile(c *gin.Context) {
	get, err := h.productRepo.GetProduct()

	if err != nil {
		domain.MessageError(c, err.Error(), http.StatusInternalServerError)
		return
	}

	if get != nil {
		var setHeader string
		var setBody string

		setHeader =
			"Product Code" + "\t" +
			"Product Name" + "\t" +
			"Product Slug" + "\t" +
			"Product Description" + "\t" +
			"Quantity" + "\t" +
			"Min Quantity" + "\t" +
			"Max Quantity" + "\t" +
			"Weight" + "\t" +
			"Volume" + "\n"

		for _, val := range *get {
			setBody +=
				val.ProductCode + "\t" +
				val.ProductName + "\t" +
				val.ProductSlug + "\t" +
				val.ProductDescription + "\t" +
				fmt.Sprintf("%v", val.Qty) + "\t" +
				fmt.Sprintf("%v", val.MinQty) + "\t" +
				fmt.Sprintf("%v", val.MaxQty) + "\t" +
				fmt.Sprintf("%v", val.Weight) + "\t" +
				fmt.Sprintf("%v", val.Volume) + "\n"
		}

		var groups = setHeader + setBody

		loc, _ := time.LoadLocation("Asia/Jakarta")
		now := time.Now().In(loc).Format("0601021504")

		fileName := "INVOICE_PRODUCT_" + now + ".txt"
		filePath := "INV_PRODUCT"
		host := config.C.FTP.FTPHOST
		port := config.C.FTP.FTPPORT
		user := config.C.FTP.FTPUSER
		pass := config.C.FTP.FTPPASS

		upload, err := service.UploadFile(groups, host, port, user, pass, fileName, filePath)

		if err != nil {
			domain.MessageError(c, err.Error(), http.StatusInternalServerError)
			return
		}

		if upload {
			domain.MessageSuccess(c, upload, "Berhasil upload file", http.StatusOK)
		}
	} else {
		domain.MessageError(c, "Data not found", http.StatusNotFound)
	}
}

// @Summary Read File
// @Description Read file from FTP
// @Accept json
// @Produce json
// @Tags FTP Controller
// @Success 200 {object} domain.Message
// @Failure 500,400,404,403 {object} domain.Message
// @Router /ReadFile [post]
func (h *ProductHandler) ReadFile(c *gin.Context) {
	fileName := "INVOICE_PRODUCT_"
	filePath := "INV_PRODUCT"
	host := config.C.FTP.FTPHOST
	port := config.C.FTP.FTPPORT
	user := config.C.FTP.FTPUSER
	pass := config.C.FTP.FTPPASS

	read, err := service.ReadFile(host, port, user, pass, fileName, filePath)

	if err != nil {
		domain.MessageError(c, err.Error(), http.StatusInternalServerError)
		return
	}

	if read != nil {
		for _, insert := range *read {
			set, err := h.productRepo.InsertProduct(insert)

			if err != nil {
				domain.MessageError(c, err.Error(), http.StatusInternalServerError)
				return
			}
			domain.MessageSuccess(c, set, "Berhasil insert product", http.StatusOK)
		}
	}
}

// @Summary Move File
// @Description Move file from FTP
// @Accept json
// @Produce json
// @Tags FTP Controller
// @Success 200 {object} domain.Message
// @Failure 500,400,404,403 {object} domain.Message
// @Router /MoveFile [post]
func (h *ProductHandler) MoveFile(c *gin.Context) {
	fileName := "INVOICE_PRODUCT_"
	filePath := "INV_PRODUCT"
	host := config.C.FTP.FTPHOST
	port := config.C.FTP.FTPPORT
	user := config.C.FTP.FTPUSER
	pass := config.C.FTP.FTPPASS

	move, err := service.MoveFile(host, port, user, pass, fileName, filePath)

	if err != nil {
		domain.MessageError(c, err.Error(), http.StatusInternalServerError)
		return
	}

	if move {
		domain.MessageSuccess(c, move, "Moving file to history successsfully", http.StatusOK)
	}
}

// @Summary Rename File
// @Description Rename file from FTP
// @Accept json
// @Produce json
// @Tags FTP Controller
// @Success 200 {object} domain.Message
// @Failure 500,400,404,403 {object} domain.Message
// @Router /RenameFile [post]
func (h *ProductHandler) RenameFile(c *gin.Context) {
	fileName := "INVOICE_PRODUCT_"
	filePath := "INV_PRODUCT"
	host := config.C.FTP.FTPHOST
	port := config.C.FTP.FTPPORT
	user := config.C.FTP.FTPUSER
	pass := config.C.FTP.FTPPASS

	rename, err := service.RenameFile(host, port, user, pass, fileName, filePath)

	if err != nil {
		domain.MessageError(c, err.Error(), http.StatusInternalServerError)
		return
	}

	if rename {
		domain.MessageSuccess(c, rename, "Rename file successfully", http.StatusOK)
	}
}

// @Summary Delete File
// @Description Delete file from FTP
// @Accept json
// @Produce json
// @Tags FTP Controller
// @Success 200 {object} domain.Message
// @Failure 500,400,404,403 {object} domain.Message
// @Router /DeleteFile [post]
func (h *ProductHandler) DeleteFile(c *gin.Context) {
	fileName := "INVOICE_PRODUCT_"
	filePath := "INV_PRODUCT"
	host := config.C.FTP.FTPHOST
	port := config.C.FTP.FTPPORT
	user := config.C.FTP.FTPUSER
	pass := config.C.FTP.FTPPASS

	del, err := service.DeleteFile(host, port, user, pass, fileName, filePath)

	if err != nil {
		domain.MessageError(c, err.Error(), http.StatusInternalServerError)
		return
	}

	if del {
		domain.MessageSuccess(c, del, "Delete file successfully", http.StatusOK)
	}
}