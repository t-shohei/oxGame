// mainパッケージ定義
package main

// パッケージ fmt, html/template, net/http をインポート
import (

	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"strings"
)

/* テンプレートの設定
ParseFiles関数（html/templateパッケージ）
    テンプレート template1.html を読み込む
    ParseFiles関数は、Template構造体のポインタを返す
Must関数（html/templateパッケージ）
    テンプレートファイルの読み込みに失敗した場合パニック発生させて
    プログラムを終了する
    Must関数は Template構造体のポインタを返す
*/
var tmpl = template.Must(template.ParseFiles("othell.html"))
type OxGame struct {
    Turn string
    List [][]string
    Win  string
}

func (game *OxGame) judge(w http.ResponseWriter, r *http.Request) {
    res := r.FormValue("res")
    if res != "" {
        val := strings.Split(res,"_")
        x,y := val[0], val[1]
        i,_ := strconv.Atoi(x)
        j,_ := strconv.Atoi(y)
        fmt.Println(x,y)
        game.List[i][j] = game.Turn
        flg := false
        lines := [][][]int  {
            {[]int{0,0},[]int{0,1},[]int{0,2}},
            {[]int{1,0},[]int{1,1},[]int{1,2}},
            {[]int{2,0},[]int{2,1},[]int{2,2}},
            {[]int{0,0},[]int{1,0},[]int{2,0}},
            {[]int{0,1},[]int{1,1},[]int{2,1}},
            {[]int{0,2},[]int{1,2},[]int{2,2}},
            {[]int{0,2},[]int{1,1},[]int{2,0}},
            {[]int{0,0},[]int{1,1},[]int{2,2}},

        }
    //勝敗ついてるか
        for i := 0; i < len(lines); i++ {
            ai, aj, bi, bj, ci,cj := lines[i][0][0], lines[i][0][1], lines[i][1][0], lines[i][1][1], lines[i][2][0], lines[i][2][1]
            if (game.List[ai][aj] == game.Turn && game.List[ai][aj] == game.List[bi][bj] && game.List[ai][aj] == game.List[ci][cj]) {
                flg = true
            }
        }
    // 勝敗ついてなかったら
        if !flg{
            if game.Turn == "o" {
                game.Turn = "x"
            }else{
                game.Turn ="o"
            }
        }       
    //勝ってたら
        if flg {
            game.Win = "勝ち"
        }
    }

    w.Header().Set("Content-Type", "text/html; charset=utf-8")
    // テンプレートとserverHTML構造体を組み合わせたHTMLをクライアントに送信
    result := tmpl.Execute(w, game)
    // クライアントへの送信にエラーがあるかを判断
    if result != nil {
        // エラーであれば panic 関数を使用して終了
        panic(result)
    }
}

func main() {
    game := &OxGame{" ", [][]string{{"", "", ""}, {"", "", ""}, {"", "", ""}}, "手番"}
    http.HandleFunc("/", game.judge)
    // Webサーバーを起動（ポート番号 8888）
    result := http.ListenAndServe(":8888", nil)
    if result != nil {
        fmt.Println(result)
    }
}