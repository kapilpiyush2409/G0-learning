package controllers

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"sync"

	"github.com/gin-gonic/gin"
)

const fileStorePath = "./filestores"

var mutex sync.Mutex

type WordCount struct {
	Word  string
	Count int
}

// ByCount implements sort.Interface based on the Count field.
type ByCount []WordCount

func (a ByCount) Len() int           { return len(a) }
func (a ByCount) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByCount) Less(i, j int) bool { return a[i].Count < a[j].Count }

// add files
func UploadMultipleTextFiles(c *gin.Context) {
	form, _ := c.MultipartForm()
	files := form.File["files"]
	fmt.Println(files)
	filePaths := []string{}
	for _, file := range files {
		// Save the file to the specified directory
		filePath := filepath.Join(fileStorePath, file.Filename)
		if err := c.SaveUploadedFile(file, filePath); err != nil {
			log.Println("Error saving file:", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to save file"})
			return
		}
		filePaths = append(filePaths, filePath)
	}
	c.JSON(http.StatusOK, gin.H{"message": "Text files uploaded successfully", "filePaths": filePaths})

}

// add file
func AddFile(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid file"})
		return
	}
	filePath := filepath.Join(fileStorePath, file.Filename)
	mutex.Lock()
	defer mutex.Unlock()

	if _, err := os.Stat(filePath); !os.IsNotExist(err) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "File already exists"})
		return
	}

	if err := c.SaveUploadedFile(file, filePath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to save file"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "File added successfully"})
}

// list files
func ListFiles(c *gin.Context) {
	files, err := os.ReadDir(fileStorePath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to list files"})
		return
	}

	var fileNames []string
	for _, file := range files {
		fileNames = append(fileNames, file.Name())
	}

	c.JSON(http.StatusOK, gin.H{"files": fileNames})
}

// remove file
func RemoveFile(c *gin.Context) {
	filename := c.Param("filename")
	filePath := filepath.Join(fileStorePath, filename)
	mutex.Lock()
	defer mutex.Unlock()

	if err := os.Remove(filePath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to remove file Or file does not exist in Store"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "File removed successfully"})
}

// update file
func UpdateFile(c *gin.Context) {
	filename := c.Param("filename")
	filePath := filepath.Join(fileStorePath, filename)
	mutex.Lock()
	defer mutex.Unlock()

	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid file"})
		return
	}

	if err := c.SaveUploadedFile(file, filePath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to update file"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "File updated successfully"})
}

func CountWordsInDirectory(c *gin.Context) {
	files, err := os.ReadDir(fileStorePath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	totalWords := 0
	for _, file := range files {
		if !file.IsDir() {
			content, err := os.ReadFile(fileStorePath + "/" + file.Name())
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			totalWords += len(strings.Fields(string(content)))
		}
	}
	c.JSON(http.StatusOK, gin.H{"word_count": totalWords})
}

func WordsFrequencyInDirectory(c *gin.Context) {
	limit, _ := c.GetQuery("limit")
	order, _ := c.GetQuery("order")
	files, err := os.ReadDir(fileStorePath)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	wordFrequency := make(map[string]int)
	for _, file := range files {
		if !file.IsDir() {
			content, err := os.ReadFile(fileStorePath + "/" + file.Name())
			if err != nil {
				c.JSON(500, gin.H{"error": err.Error()})
				return
			}

			scanner := bufio.NewScanner(strings.NewReader(string(content)))
			scanner.Split(bufio.ScanWords)
			for scanner.Scan() {
				word := scanner.Text()
				wordFrequency[word]++
			}
		}
	}

	var wordCounts []WordCount
	for word, count := range wordFrequency {
		wordCounts = append(wordCounts, WordCount{Word: word, Count: count})
	}
	// Sort words by frequency
	if order == "dsc" {
		sort.Sort(sort.Reverse(ByCount(wordCounts)))
	} else {
		sort.Sort(ByCount(wordCounts))
	}

	// Apply limit
	limitInt := 10 // default value
	if limit != "" {
		limitInt, err = strconv.Atoi(limit)
		if err != nil {
			limitInt = 10
		}
	}
	if limitInt < len(wordCounts) {
		wordCounts = wordCounts[:limitInt]
	}

	c.JSON(200, wordCounts)
}
