package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/fatih/color"
)

type item struct {
	Task        string
	Done        bool
	CreatedAt   time.Time
	CompletedAt time.Time
}

type List []item

func (l *List) Add(task string) {
	t := item{
		Task:        task,
		Done:        false,
		CreatedAt:   time.Now(),
		CompletedAt: time.Time{},
	}

	*l = append(*l, t)

}

func (l *List) Complete(i int) error {
	ls := *l
	if i <= 0 || i > len(ls) {
		return fmt.Errorf("item %d does not exist", i)
	}

	ls[i-1].Done = true
	ls[i-1].CompletedAt = time.Now()

	return nil
}

// Delete removes an element from the List at the specified index.
//
// It takes an integer parameter `i` which represents the index of the element to be deleted.
// The function returns an error if the index is out of range.
// It returns nil if the element is successfully deleted.
func (l *List) Delete(i int) error {
	ls := *l
	if i <= 0 || i > len(ls) {
		return fmt.Errorf("item %d does not exist", i)
	}

	*l = append(ls[:i-1], ls[i:]...)

	return nil

}

// Save saves the List struct to a file in JSON format.
//
// filename: the name of the file to save the List to.
// error: an error if there was an issue saving the List.
func (l *List) Save(filename string) error {
	js, err := json.Marshal(l)
	if err != nil {
		return err
	}

	return os.WriteFile(filename, js, 0644)
}

func (l *List) Get(filename string) error {
	file, err := os.ReadFile(filename)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil
		}
		return err
	}
	if len(file) == 0 {
		return nil
	}
	return json.Unmarshal(file, l)
}

func (l *List) String() string {
	formatted := ""
	c := color.New()
	for k, t := range *l {
		prefix := "    "
		if t.Done {
			prefix = "[x] "
			c.Add(color.FgGreen)
			formatted += c.Sprintf("%s%d: %s\n", prefix, k+1, t.Task)
		} else {
			c.Add(color.FgRed)
			formatted += c.Sprintf("%s%d: %s\n", prefix, k+1, t.Task)
		}
	}
	return formatted
}
