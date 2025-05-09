package main

import (
    "bufio"
    "fmt"
    "math/rand"
    "os"
    "strconv"
    "strings"
    "time"
)

const (
    colorReset  = "\033[0m"
    colorRed    = "\033[31m"
    colorGreen  = "\033[32m"
    colorYellow = "\033[33m"
    colorCyan   = "\033[36m"
    colorWhite  = "\033[97m"
)

func clearScreen() {
    fmt.Print("\033[2J\033[H") // ANSI: очистка экрана и курсор в начало
}

func chooseLevel() (int, string) {
    fmt.Println(colorCyan + "Choose difficulty:" + colorReset)
    fmt.Println("1 - Easy (1-10)")
    fmt.Println("2 - Medium (1-50)")
    fmt.Println("3 - Hard (1-100)")

    reader := bufio.NewReader(os.Stdin)
    fmt.Print(colorYellow + "Enter level: " + colorReset)
    input, _ := reader.ReadString('\n')
    level := strings.TrimSpace(input)

    switch level {
    case "1":
        return 10, "Easy"
    case "2":
        return 50, "Medium"
    default:
        return 100, "Hard"
    }
}

func main() {
    rand.Seed(time.Now().UnixNano())
    clearScreen()

    max, levelName := chooseLevel()
    secret := rand.Intn(max) + 1
    attempts := 0

    reader := bufio.NewReader(os.Stdin)

    for {
        clearScreen()
        fmt.Println(colorCyan + "🔢 GUESS THE NUMBER" + colorReset)
        fmt.Printf("🎮 Difficulty: %s (%s1-%d%s)\n", levelName, colorYellow, max, colorReset)
        fmt.Printf("🔁 Attempts: %s%d%s\n", colorWhite, attempts, colorReset)
        fmt.Print(colorGreen + "👉 Enter your guess: " + colorReset)

        input, _ := reader.ReadString('\n')
        guessStr := strings.TrimSpace(input)
        guess, err := strconv.Atoi(guessStr)

        if err != nil {
            fmt.Println(colorRed + "Invalid number! Try again." + colorReset)
            time.Sleep(1 * time.Second)
            continue
        }

        attempts++

        if guess < secret {
            fmt.Println(colorYellow + "📉 Too low!" + colorReset)
        } else if guess > secret {
            fmt.Println(colorYellow + "📈 Too high!" + colorReset)
        } else {
            fmt.Printf(colorGreen+"🎉 Correct! You guessed it in %d attempts!\n"+colorReset, attempts)
            break
        }

        fmt.Println(colorWhite + "Press Enter to try again..." + colorReset)
        reader.ReadString('\n')
    }
}
