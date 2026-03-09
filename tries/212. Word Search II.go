package main

import "fmt"

// TrieNode mewakili satu kotak huruf di dalam pohon Trie.
type TrieNode struct {
	// children adalah array berisi 26 pointer ke node anak (a-z).
	// Jika children[0] tidak nil, artinya ada jalur huruf 'a'.
	children [26]*TrieNode
	// word menyimpan kata lengkap jika node ini adalah akhir dari sebuah kata.
	// Jika kosong (""), artinya ini bukan akhir kata.
	// (Ini menggantikan boolean 'isEnd').
	word string
}

// buildTrie memasukkan semua kata dari slice 'words' ke dalam struktur Trie.
func buildTrie(words []string) *TrieNode {
	// Buat node Root kosong sebagai titik awal pohon.
	root := &TrieNode{}
	
	// Iterasi setiap kata dalam daftar.
	for _, w := range words {
		// Mulai dari root untuk setiap kata baru.
		curr := root
		
		// Iterasi setiap karakter dalam kata.
		for i := 0; i < len(w); i++ {
			// Hitung index array (0-25) berdasarkan huruf ('a'=0, 'b'=1, dst).
			idx := w[i] - 'a'
			
			// Jika jalur untuk huruf ini belum ada (nil)...
			if curr.children[idx] == nil {
				// ...buat node baru untuk jalur tersebut.
				curr.children[idx] = &TrieNode{}
			}
			
			// Pindah pointer 'curr' ke node anak tersebut.
			curr = curr.children[idx]
		}
		// Setelah semua huruf dimasukkan, simpan kata lengkap di node terakhir.
		curr.word = w
	}
	// Kembalikan root dari Trie yang sudah jadi.
	return root
}

// findWords adalah fungsi utama untuk mencari semua kata di papan.
func findWords(board [][]byte, words []string) []string {
	// 1. Bangun Trie dari semua kata yang ingin dicari.
	root := buildTrie(words)
	// res akan menampung kata-kata yang berhasil ditemukan.
	res := []string{}
	
	// Ambil jumlah baris dan kolom papan.
	rows := len(board)
	cols := len(board[0])
	
	// 2. Iterasi setiap kotak di papan (R = baris, C = kolom).
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			// Mulai pencarian DFS dari kotak (r, c) menggunakan Trie Root.
			dfs(r, c, board, root, &res)
		}
	}
	// Kembalikan daftar kata yang ditemukan.
	return res
}

// dfs adalah fungsi rekursif untuk menjelajahi papan (Backtracking).
// r: baris sekarang, c: kolom sekarang, node: node Trie saat ini.
func dfs(r, c int, board [][]byte, node *TrieNode, res *[]string) {
	// --- A. PENGECEKAN BATAS & KONDISI AWAL ---
	
	// 1. Cek apakah koordinat (r, c) di luar batas papan.
	if r < 0 || r >= len(board) || c < 0 || c >= len(board[0]) {
		return
	}
	
	// 2. Ambil huruf di kotak sekarang.
	char := board[r][c]
	
	// 3. Cek kondisi valid:
	// - Jika char == '#', artinya kotak ini sudah dikunjungi dalam jalur kata yang sama.
	// - Jika node.children[char-'a'] == nil, artinya tidak ada kata di Trie yang mulai dengan huruf ini.
	if char == '#' || node.children[char-'a'] == nil {
		return
	}
	
	// --- B. PINDAH KE DALAM TRIE ---
	
	// Karena huruf valid, kita pindah ke node anak yang sesuai di Trie.
	node = node.children[char-'a']
	
	// --- C. CEK JIKA KATA DITEMUKAN ---
	
	// Jika node ini menyimpan kata (tidak kosong)...
	if node.word != "" {
		// ...masukkan kata tersebut ke dalam hasil.
		*res = append(*res, node.word)
		// HAPUS kata di Trie (diubah jadi kosong).
		// Ini trik penting agar kata yang sama tidak ditemukan dua kali (duplikat).
		node.word = ""
	}
	
	// --- D. MARKING (TANDAI SUDAH DIKUNJUNGI) ---
	
	// Ubah huruf di papan sementara menjadi '#' agar tidak dipakai dua kali
	// dalam satu formasi kata yang sedang dicari (menghindari looping).
	board[r][c] = '#'
	
	// --- E. REKURSI (JALAN KE 4 ARAH) ---
	
	// Panggil DFS untuk kotak Tetangga: Atas, Bawah, Kiri, Kanan.
	dfs(r+1, c, board, node, res) // Bawah
	dfs(r-1, c, board, node, res) // Atas
	dfs(r, c+1, board, node, res) // Kanan
	dfs(r, c-1, board, node, res) // Kiri
	
	// --- F. BACKTRACKING (KEMBALIKAN HURUF) ---
	
	// Setelah selesai menjelajahi 4 arah dari kotak ini (rekursi selesai),
	// kita harus mengembalikan huruf aslinya ('#' diubah balik jadi 'char').
	// Ini supaya kotak ini bisa dipakai kembali untuk jalur pencarian kata lain.
	board[r][c] = char
}

func main() {
	// Contoh penggunaan:
	board := [][]byte{
		{'o', 'a', 'a', 'n'},
		{'e', 't', 'a', 'e'},
		{'i', 'h', 'k', 'r'},
		{'i', 'f', 'l', 'v'},
	}
	words := []string{"oath", "pea", "eat", "rain"}
	
	result := findWords(board, words)
	fmt.Println("Kata yang ditemukan:", result) 
	// Output: Kata yang ditemukan: [oath eat] (urutan bisa berbeda)
}