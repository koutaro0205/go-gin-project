package example

import "testing"

// NOTE: テスト対象の関数名は"TestXXX"とする
func TestPerson_Greeting(t *testing.T) {
	if testing.Short() {
		t.SkipNow() // --shortオプションで、テストをスキップ
	}

	type fields struct {
		Name string
		Age  int
	}

	type args struct {
		text string
	}

	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{"テスト1", fields{"山田", 22}, args{"こんにちは"}, "山田(22)さん, こんにちは"},
		{"テスト2", fields{"佐藤", 26}, args{"はじめまして"}, "佐藤(26)さん, はじめまして"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// ログを表示
			t.Logf("t.Logf  %s \n", tt.name)

			person := &Person{
				Name: tt.fields.Name,
				Age:  tt.fields.Age,
			}

			if got := person.Greeting(tt.args.text); got != tt.want {
				// NOTE: t.Errorf→テストが失敗しても継続する
				// NOTE: 継続したくない場合は、t.FailNow()もしくはt.Fatalf()
				t.Errorf("Greeting() = %v, want %v", got, tt.want)
			}
		})
	}
}
