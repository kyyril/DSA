func isValid(s string) bool {
    stack := []byte{}
    pairs := map[byte]byte{
        ')': '(',
        ']': '[',
        '}': '{',
    }

    for i := 0; i < len(s); i++ {
        c := s[i]

        if c == '(' || c == '[' || c == '{' {
            stack = append(stack, c)
        } else {
            if len(stack) == 0 {
                return false
            }
            top := stack[len(stack)-1]
            stack = stack[:len(stack)-1]

            if top != pairs[c] {
                return false
            }
        }
    }

    return len(stack) == 0
}
