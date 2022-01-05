#include <string.h>
#include <ctype.h>
#include <stdlib.h>
#include <stdbool.h>

typedef struct stack {
    int *data;
    int top;
} stack;

void InitStack(stack *st, int cap) {
    st->data = malloc(sizeof(int) * cap);
    st->top = 0;
}

void Push(stack *st, int val) {
    st->data[st->top] = val;
    st->top = st->top + 1;
}

int Pop(stack *st) {
    st->top = st->top - 1;
    return st->data[st->top];
}

int Top(stack *st) {
    return st->data[st->top - 1];
}

bool IsEmpty(stack *st) {
    return st->top == 0;
}

int priority(int opt) {
    if (opt == '+' || opt == '-') {
        return 1;
    }
    return 2;
}

void calc(stack *nums, stack *opts) {
    int b = Pop(nums);
    int a = Pop(nums);
    int opt = Pop(opts);
    int ans = 0;
    if (opt == '+') ans = a + b;
    else if (opt == '-') ans = a - b;
    else if (opt == '*') ans = a * b;
    else if (opt == '/') ans = a / b;
    Push(nums, ans);
}

int calculate(char *s) {
    int n = strlen(s), num = 0;
    bool flag = false;
    stack nums, opts;
    InitStack(&nums, n);
    InitStack(&opts, n);
    for (int i = 0; i < n; i++) {
        if (isdigit(s[i])) {
            num *= 10;
            num += s[i] - '0';
            flag = true;
        };

        if (flag && (!isdigit(s[i]) || i == n - 1)) {
            Push(&nums, num);
            num = 0;
            flag = false;
        }

        if (!isdigit(s[i]) && s[i] != ' ') {
            while (!IsEmpty(&opts)
                   && priority(Top(&opts))
                      >= priority(s[i])) {
                calc(&nums, &opts);
            }
            Push(&opts, s[i]);
        }
    }
    while (!IsEmpty(&opts)) calc(&nums, &opts);
    return Top(&nums);
}