package format

import "fmt"

func FormatSec(sec int64) string {
    min := sec / 60
    h := min / 60
    s := sec - min * 60
    min = min - h * 60
    return fmt.Sprintf("%02d:%02d:%02d", h, min, s)
}

