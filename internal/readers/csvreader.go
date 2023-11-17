package readers

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"strings"

	"github.com/jawee/scim-client-go/internal/models"
)

func ReadFile(in io.Reader) ([]models.User, error) {

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

        log.Printf("%v\n", rec)

        user := models.User {
            Id:  rec[indexes[0]],
            UserName: rec[indexes[1]],
            Email: rec[indexes[1]],
            FirstName: rec[indexes[2]],
            LastName: rec[indexes[3]],
            Department: rec[indexes[4]],
            PhoneNumber: rec[indexes[5]],
        }

        users = append(users, user)
    }

    return users, nil
}

func columnNotFoundError(column string) error {
    return fmt.Errorf("Column '%s' not found.", column)
}

func getIndexes(headerRow []string) ([]int, error) {
    indexes := make([]int, 0)

    idx := indexOf(headerRow, "Id")
    if idx == -1 {
        return nil, columnNotFoundError("Id")
    }
    indexes = append(indexes, idx)
    idx = indexOf(headerRow, "Email")
    if idx == -1 {
        return nil, columnNotFoundError("Email")
    }
    indexes = append(indexes, idx)

    idx = indexOf(headerRow, "FirstName")
    if idx == -1 {
        return nil, columnNotFoundError("FirstName")
    }
    indexes = append(indexes, idx)

    idx = indexOf(headerRow, "LastName")
    if idx == -1 {
        return nil, columnNotFoundError("LastName")
    }
    indexes = append(indexes, idx)

    idx = indexOf(headerRow, "Department")
    if idx == -1 {
        return nil, columnNotFoundError("Department")
    }
    indexes = append(indexes, idx)

    idx = indexOf(headerRow, "MobilePhone")
    if idx == -1 {
        return nil, columnNotFoundError("MobilePhone")
    }
    indexes = append(indexes, idx)

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
