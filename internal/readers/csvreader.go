package readers

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/jawee/scim-client-go/internal/configuration"
	"github.com/jawee/scim-client-go/internal/models"
)

type CsvReader struct {}

func (c *CsvReader) GetUsers(config configuration.FileReaderConfig) ([]models.User, error) {
    file, err := os.Open(config.FilePath)
    if err != nil {
        return nil, err
    }

    users, err := readFile(file)
    if err != nil {
        return nil, err
    }

    return users, nil;
}

func readFile(in io.Reader) ([]models.User, error) {
    users := make([]models.User, 0)

    csvReader := csv.NewReader(in)
    headerRow, err := csvReader.Read()
    if err != nil {
        return nil, err
    }

    indexes, err := getIndexes(headerRow)
    if err != nil {
        return nil, err
    }

    for {
        rec, err := csvReader.Read()
        if err == io.EOF {
            break
        }
        if err != nil {
            return nil, err
        }

        for i, v := range rec {
            rec[i] = strings.TrimSpace(v)
        }

        user := models.User {
            Id:  rec[indexes[ID]],
            UserName: rec[indexes[EMAIL]],
            Email: rec[indexes[EMAIL]],
            FirstName: rec[indexes[FIRSTNAME]],
            LastName: rec[indexes[LASTNAME]],
            Department: rec[indexes[DEPARTMENT]],
            PhoneNumber: rec[indexes[PHONENUMBER]],
        }

        users = append(users, user)
    }

    return users, nil
}

var (
    ID = 0
    EMAIL = 1
    FIRSTNAME = 2
    LASTNAME = 3
    DEPARTMENT = 4
    PHONENUMBER = 5
)

func columnNotFoundError(column string) error {
    return fmt.Errorf("Column '%s' not found.", column)
}

func getIndexForColumn(row []string, column string) (int, error) {
    idx := indexOf(row, column)
    if idx == -1 {
        return -1, columnNotFoundError("Id")
    }

    return idx, nil
}

func getIndexes(headerRow []string) ([]int, error) {
    indexes := make([]int, 0)

    columns := []string { "Id", "Email", "FirstName", "LastName", "Department", "MobilePhone" }

    for _, v := range columns {
        idx, err := getIndexForColumn(headerRow, v)
        if err != nil {
            return nil, err
        }
        indexes = append(indexes, idx)
    }

    return indexes, nil
}

func indexOf(haystack []string, needle string) int {
    for i, v := range haystack {
        if v == needle {
            return i
        }
    }
    return -1
}
