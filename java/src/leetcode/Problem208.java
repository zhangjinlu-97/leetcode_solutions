package leetcode;

import java.util.List;

public class Problem208 {

}

class Trie {

    Node root;

    static class Node {
        int pass;
        int end;
        Node[] nextNodes;

        public Node() {
            nextNodes = new Node[26];
        }
    }

    public Trie() {
        root = new Node();
    }

    public void insert(String word) {
        root.pass++;
        Node cur = root;
        for (int i = 0; i < word.length(); i++) {
            int path = word.charAt(i) - 'a';
            if (cur.nextNodes[path] == null) {
                cur.nextNodes[path] = new Node();
            }
            cur = cur.nextNodes[path];
            cur.pass++;
        }
        cur.end++;
    }

    public boolean search(String word) {
        Node cur = root;
        for (int i = 0; i < word.length(); i++) {
            int path = word.charAt(i) - 'a';
            if (cur.nextNodes[path] == null) {
                return false;
            }
            cur = cur.nextNodes[path];
        }
        return cur.end > 0;
    }

    public boolean startsWith(String prefix) {
        Node cur = root;
        for (int i = 0; i < prefix.length(); i++) {
            int path = prefix.charAt(i) - 'a';
            if (cur.nextNodes[path] == null) {
                return false;
            }
            cur = cur.nextNodes[path];
        }
        return true;
    }
}
