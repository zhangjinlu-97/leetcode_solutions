#include <stdio.h>
#include "datastructure.h"

void process(struct TreeNode *root, int *k, int *res) {
    if (root == NULL || *k < 0) return;
    process(root->left, k, res);
    if (--(*k) == 0) *res = root->val;
    process(root->right, k, res);
}

int kthSmallest(struct TreeNode *root, int k) {
    int res;
    process(root, &k, &res);
    return res;
}

