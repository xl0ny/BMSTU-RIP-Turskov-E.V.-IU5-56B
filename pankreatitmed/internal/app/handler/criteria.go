package handler

import (
	"fmt"
	"net/http"
	"pankreatitmed/internal/app/dto"
	"pankreatitmed/internal/app/dto/request"
	"pankreatitmed/internal/app/dto/response"
	"pankreatitmed/internal/app/mapper"
	"strconv"

	"github.com/gin-gonic/gin"

	"log"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

//const demoUserID uint = 1 // пока без авторизации

// CriteriaList GET /criteria
func (h *Handler) CriteriaList(c *gin.Context) {
	var query request.GetCriteria
	if err := c.ShouldBindQuery(&query); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	list, err := h.svcs.Criteria.List(query.Query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	items := mapper.CriterionsToSendCrtierions(list)
	res := dto.List[response.SendCriterion]{Items: items}

	c.JSON(http.StatusOK, res)

}

// CriteriaGet GET /criteria
func (h *Handler) CriteriaGet(c *gin.Context) {
	var id request.GetCriterion
	if err := c.ShouldBindUri(&id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	println("id", id.ID)
	criterion, err := h.svcs.Criteria.Get(id.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	crit := mapper.CritertionToSendCriterionLink(criterion)
	c.JSON(http.StatusOK, crit)
}

func (h *Handler) CriteriaCreate(c *gin.Context) {
	var criterion request.CreateCriterion
	if err := c.ShouldBindJSON(&criterion); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		println(1)
		return
	}
	crit, err := mapper.CreateCriterionToCriterion(criterion)
	if err != nil {
		println(2)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	if err := h.svcs.Criteria.Create(&crit); err != nil {
		println(3)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	c.Status(http.StatusOK)
}

// CriteriaUpdate TODO разобраться почему не кидает ошибку
func (h *Handler) CriteriaUpdate(c *gin.Context) {
	var id request.GetCriterion
	var criterion request.UpdateCriterion
	if err := c.ShouldBindUri(&id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := c.ShouldBindJSON(&criterion); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.svcs.Criteria.Update(id.ID, &criterion); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	c.Status(http.StatusOK)
}

func (h *Handler) CriteriaDelete(c *gin.Context) {
	var id request.GetCriterion
	if err := c.ShouldBindUri(&id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.svcs.Criteria.Delete(id.ID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	client := connectMinio() // из предыдущего примера
	if err := h.svcs.Criteria.DeleteImage(client, id.ID, c); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	c.Status(http.StatusOK)
}

// TODO настроить нормализацию последовательности БД
func (h *Handler) AddCriteriaToDraft(c *gin.Context) {
	var id request.GetCriterion
	if err := c.ShouldBindUri(&id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.svcs.Criteria.ToDraft(id.ID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
}

func (h *Handler) UploadCriterionImage(c *gin.Context) {
	// ID услуги из URL
	var id request.GetCriterion
	if err := c.ShouldBindUri(&id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fileHeader, err := c.FormFile("image")
	if err != nil {
		c.JSON(400, gin.H{"error": "image is required"})
		return
	}

	// Открываем файл
	file, err := fileHeader.Open()
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	defer file.Close()

	// Загружаем в MinIO
	client := connectMinio() // из предыдущего примера
	bucket := "services-images"
	objectName := fmt.Sprintf("service_%s/%s", strconv.Itoa(int(id.ID)), fileHeader.Filename)

	_, err = client.PutObject(
		c, bucket, objectName, file, fileHeader.Size,
		minio.PutObjectOptions{ContentType: fileHeader.Header.Get("Content-Type")},
	)
	if err != nil {
		fmt.Println("ТУТ")
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	imgname := fmt.Sprintf("http://localhost:9000/%s/%s", bucket, objectName)
	crit := request.UpdateCriterion{ImageURL: &imgname}
	fmt.Println(crit.ImageURL)
	if err := h.svcs.Criteria.DeleteImage(client, id.ID, c); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	if err := h.svcs.Criteria.Update(id.ID, &crit); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(200, gin.H{
		"status": "ok",
		"url":    fmt.Sprintf("http://localhost:9000/%s/%s", bucket, objectName),
	})
}

func connectMinio() *minio.Client {
	endpoint := "localhost:9000" // адрес контейнера
	accessKey := "minio"         // MINIO_ROOT_USER
	secretKey := "minio124"      // MINIO_ROOT_PASSWORD
	useSSL := false

	client, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKey, secretKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		log.Fatalln(err)
	}

	return client
}
