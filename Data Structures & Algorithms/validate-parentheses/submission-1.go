func isValid(s string) bool {
    pairs := map[rune]rune{
        ')': '(',
        ']': '[',
        '}': '{',
    }
    stack := []rune{}
    for _, r := range s {
        switch r {
        case '(', '[', '{':
            stack = append(stack, r)

        case ')', ']', '}':
            if len(stack) == 0 {
                return false
            }

            top := stack[len(stack)-1]
            stack = stack[:len(stack)-1]

            if top != pairs[r] {
                return false
            }
        }
    }

    return len(stack) == 0
}

