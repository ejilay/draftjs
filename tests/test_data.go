package tests

type TestTable struct {
	Name     string
	State    string
	Expected string
}

func GetTestsTable() []TestTable {
	return []TestTable{
		{
			Name:     "1",
			State:    `{"entityMap":{},"blocks":[{"key":"4g603","text":"dasdasdasdsadsaывфвыфв","type":"unstyled","depth":0,"inlineStyleRanges":[{"offset":0,"length":22,"style":"BOLD"}],"entityRanges":[],"data":{}}]}`,
			Expected: `<p><strong>dasdasdasdsadsaывфвыфв</strong></p>`,
		},
		{
			Name:     "2",
			State:    `{"entityMap":{"0":{"type":"LINK","mutability":"MUTABLE","data":{"url":"ya.ru"}}},"blocks":[{"key":"4g603","text":"dasdasdasdsadsaывфвыфв","type":"unstyled","depth":0,"inlineStyleRanges":[],"entityRanges":[{"offset":12,"length":2,"key":0}],"data":{}}]}`,
			Expected: `<p>dasdasdasdsa<a href="ya.ru" target="_blank">ds</a>aывфвыфв</p>`,
		},
		{
			Name:     "3",
			State:    `{"entityMap":{},"blocks":[{"key":"18u09","text":"\"Война и мир\" на экранах Би-Би-Си","type":"unstyled","depth":0,"inlineStyleRanges":[],"entityRanges":[],"data":{}}]}`,
			Expected: `<p>&#34;Война и мир&#34; на экранах Би-Би-Си</p>`,
		},
		{
			Name:     "4",
			State:    `{"entityMap":{},"blocks":[{"key":"ek0ec","text":"Яндекс.Метро — интерактивная карта метро Москвы с расчётом времени и прокладкой маршрутов с учётом данных о закрытии станций и вес","type":"unstyled","depth":0,"inlineStyleRanges":[{"offset":7,"length":5,"style":"ITALIC"},{"offset":35,"length":5,"style":"ITALIC"}],"entityRanges":[],"data":{}}]}`,
			Expected: `<p>Яндекс.<em>Метро</em> — интерактивная карта <em>метро</em> Москвы с расчётом времени и прокладкой маршрутов с учётом данных о закрытии станций и вес</p>`,
		},
		{
			Name:     "5",
			State:    `{"entityMap":{},"blocks":[{"key":"fdcqk","text":"Тест 11 Тест 11 Тест 11  Тест 11  Тест 11  Тест 11  Тест 11 ","type":"unstyled","depth":0,"inlineStyleRanges":[],"entityRanges":[],"data":{}}]}`,
			Expected: `<p>Тест 11 Тест 11 Тест 11  Тест 11  Тест 11  Тест 11  Тест 11 </p>`,
		},
		{
			Name:     "6",
			State:    `{"entityMap":{},"blocks":[{"key":"ao1cv","text":"Парам пам пам пам пам ","type":"unstyled","depth":0,"inlineStyleRanges":[],"entityRanges":[],"data":{}}]}`,
			Expected: `<p>Парам пам пам пам пам </p>`,
		},
		{
			Name:     "7",
			State:    `{"entityMap":{},"blocks":[{"key":"8lobd","text":"a","type":"unordered-list-item","depth":0,"inlineStyleRanges":[],"entityRanges":[],"data":{}},{"key":"7pppu","text":"b","type":"unordered-list-item","depth":0,"inlineStyleRanges":[],"entityRanges":[],"data":{}},{"key":"fn3d3","text":"c","type":"unordered-list-item","depth":0,"inlineStyleRanges":[],"entityRanges":[],"data":{}},{"key":"50oms","text":"1","type":"ordered-list-item","depth":1,"inlineStyleRanges":[],"entityRanges":[],"data":{}},{"key":"3rh3g","text":"2","type":"ordered-list-item","depth":1,"inlineStyleRanges":[],"entityRanges":[],"data":{}},{"key":"7mqrj","text":"3","type":"ordered-list-item","depth":1,"inlineStyleRanges":[],"entityRanges":[],"data":{}}]}`,
			Expected: "<ul><li>a</li><li>b</li><li>c<ol><li>1</li><li>2</li><li>3</li></ol></li></ul>",
		},
		{
			Name:     "8",
			State:    `{"entityMap":{},"blocks":[{"key":"citfp","text":"pp1","type":"ordered-list-item","depth":0,"inlineStyleRanges":[],"entityRanges":[],"data":{}},{"key":"1jje0","text":"rgh","type":"ordered-list-item","depth":0,"inlineStyleRanges":[],"entityRanges":[],"data":{}},{"key":"b6b9k","text":"pp4","type":"ordered-list-item","depth":0,"inlineStyleRanges":[],"entityRanges":[],"data":{}},{"key":"7t6hi","text":"pp","type":"ordered-list-item","depth":1,"inlineStyleRanges":[],"entityRanges":[],"data":{}},{"key":"1a0r1","text":"dwf","type":"ordered-list-item","depth":2,"inlineStyleRanges":[{"offset":0,"length":3,"style":"ITALIC"}],"entityRanges":[],"data":{}},{"key":"558bm","text":"rge","type":"ordered-list-item","depth":2,"inlineStyleRanges":[],"entityRanges":[],"data":{}},{"key":"a0d0g","text":"wdf","type":"ordered-list-item","depth":2,"inlineStyleRanges":[{"offset":0,"length":3,"style":"BOLD"}],"entityRanges":[],"data":{}},{"key":"fkok6","text":"efg","type":"ordered-list-item","depth":1,"inlineStyleRanges":[],"entityRanges":[],"data":{}},{"key":"81qon","text":"bhn","type":"ordered-list-item","depth":2,"inlineStyleRanges":[],"entityRanges":[],"data":{}},{"key":"1imp3","text":"wefg","type":"ordered-list-item","depth":0,"inlineStyleRanges":[],"entityRanges":[],"data":{}},{"key":"drgoj","text":"pp2","type":"ordered-list-item","depth":1,"inlineStyleRanges":[{"offset":0,"length":3,"style":"BOLD"}],"entityRanges":[],"data":{}},{"key":"hd28","text":"pp3","type":"unordered-list-item","depth":2,"inlineStyleRanges":[],"entityRanges":[],"data":{}},{"key":"8g8r7","text":"asd","type":"unordered-list-item","depth":2,"inlineStyleRanges":[],"entityRanges":[],"data":{}}]}`,
			Expected: "<ol><li>pp1</li><li>rgh</li><li>pp4<ol><li>pp<ol><li><em>dwf</em></li><li>rge</li><li><strong>wdf</strong></li></ol></li><li>efg<ol><li>bhn</li></ol></li></ol></li><li>wefg<ol><li><strong>pp2</strong><ul><li>pp3</li><li>asd</li></ul></li></ol></li></ol>",
		},
		{
			Name:     "9",
			State:    `{"entityMap":{},"blocks":[{"key":"4g603","text":"H2O","type":"unstyled","depth":0,"inlineStyleRanges":[{"offset":1,"length":1,"style":"SUBSCRIPT"}],"entityRanges":[],"data":{}}]}`,
			Expected: "<p>H<sub>2</sub>O</p>",
		},
		{
			Name:     "10",
			State:    `{"entityMap":{},"blocks":[{"key":"4g603","text":"210 = 1024","type":"unstyled","depth":0,"inlineStyleRanges":[{"offset":1,"length":2,"style":"SUPERSCRIPT"}],"entityRanges":[],"data":{}}]}`,
			Expected: "<p>2<sup>10</sup> = 1024</p>",
		},
		{
			Name:     "11",
			State:    `{"blocks":[{"type":"unstyled","inlineStyleRanges":[],"entityRanges":[],"data":{},"text":"Описание к изображению","depth":0}],"entityMap":{}}`,
			Expected: "<p>Описание к изображению</p>",
		},
		{
			Name:     "12",
			State:    `{"blocks":[{"type":"unstyled","inlineStyleRanges":[],"entityRanges":[],"data":{},"text":"Я прочитала книгу на английском, как только она появилась. На русском читать не буду, мне не нравится этот перевод.","depth":0}],"entityMap":{}}`,
			Expected: "<p>Я прочитала книгу на английском, как только она появилась. На русском читать не буду, мне не нравится этот перевод.</p>",
		},
		{
			Name:     "13",
			State:    `{"blocks":[{"type":"unstyled","inlineStyleRanges":[],"entityRanges":[],"data":{},"text":"Наталья Водянова, российская супермодель","depth":0}],"entityMap":{}}`,
			Expected: "<p>Наталья Водянова, российская супермодель</p>",
		},
		{
			Name:     "14",
			State:    `{"blocks":[{"type":"unstyled","inlineStyleRanges":[],"entityRanges":[],"data":{},"text":"Some Words Here aaaaaaaa","depth":0}],"entityMap":{}}`,
			Expected: "<p>Some Words Here aaaaaaaa</p>",
		},
		{
			Name:     "15",
			State:    `{"blocks":[{"type":"unstyled","inlineStyleRanges":[],"entityRanges":[],"data":{},"text":"Не может быть идеального продукта. Если ты уверен, что создал идеальный продукт или идеальное СМИ, то тебе пора на пенсию.","depth":0}],"entityMap":{}}`,
			Expected: "<p>Не может быть идеального продукта. Если ты уверен, что создал идеальный продукт или идеальное СМИ, то тебе пора на пенсию.</p>",
		},
		{
			Name:     "16",
			State:    `{"blocks":[{"type":"unstyled","inlineStyleRanges":[],"entityRanges":[],"data":{},"text":"","depth":0}],"entityMap":{}}`,
			Expected: "<p></p>",
		},
		{
			Name:     "17",
			State:    `{"blocks":[{"type":"unstyled","inlineStyleRanges":[],"entityRanges":[],"data":{},"text":"Текст карточки с картинкой","depth":0}],"entityMap":{}}`,
			Expected: "<p>Текст карточки с картинкой</p>",
		},
		{
			Name:     "18",
			State:    `{"blocks":[{"type":"unstyled","inlineStyleRanges":[],"entityRanges":[],"data":{},"text":"Абзац текста в карточке","depth":0}],"entityMap":{}}`,
			Expected: "<p>Абзац текста в карточке</p>",
		},
		{
			Name:     "19",
			State:    `{"blocks":[{"type":"unstyled","inlineStyleRanges":[],"entityRanges":[],"data":{},"text":"Заголовок карточки","depth":0}],"entityMap":{}}`,
			Expected: "<p>Заголовок карточки</p>",
		},
		{
			Name:     "20",
			State:    `{"blocks":[{"type":"unstyled","inlineStyleRanges":[],"entityRanges":[],"data":{},"text":"Some Words Here aaaaaaaa","depth":0}],"entityMap":{}}`,
			Expected: "<p>Some Words Here aaaaaaaa</p>",
		},
		{
			Name:     "21",
			State:    `{"blocks":[{"type":"unstyled","inlineStyleRanges":[],"entityRanges":[],"data":{},"text":"Не может быть идеального продукта. Если ты уверен, что создал идеальный продукт или идеальное СМИ, то тебе пора на пенсию.","depth":0}],"entityMap":{}}`,
			Expected: "<p>Не может быть идеального продукта. Если ты уверен, что создал идеальный продукт или идеальное СМИ, то тебе пора на пенсию.</p>",
		},
		{
			Name:     "22",
			State:    `{"blocks":[{"type":"ordered-list-item","inlineStyleRanges":[],"entityRanges":[],"data":{},"text":"1 элемент нумерованного списка","depth":0},{"type":"ordered-list-item","inlineStyleRanges":[],"entityRanges":[],"data":{},"text":"2 элемент нумерованного списка","depth":0}],"entityMap":{}}`,
			Expected: "<ol><li>1 элемент нумерованного списка</li><li>2 элемент нумерованного списка</li></ol>",
		},
		{
			Name:     "23",
			State:    `{"blocks":[{"type":"unordered-list-item","inlineStyleRanges":[],"entityRanges":[],"data":{},"text":"1 элемент ненумерованного списка","depth":0},{"type":"unordered-list-item","inlineStyleRanges":[],"entityRanges":[],"data":{},"text":"2 элемент ненумерованного списка","depth":0}],"entityMap":{}}`,
			Expected: "<ul><li>1 элемент ненумерованного списка</li><li>2 элемент ненумерованного списка</li></ul>",
		},
		{
			Name:     "24",
			State:    `{"blocks":[{"type":"unstyled","inlineStyleRanges":[{"length":5,"offset":0,"style":"ITALIC"},{"length":6,"offset":6,"style":"BOLD"},{"length":19,"offset":15,"style":"UNDERLINE"}],"entityRanges":[],"data":{},"text":"Абзац текста с рич-форматированием","depth":0}],"entityMap":{}}`,
			Expected: "<p><em>Абзац</em> <strong>текста</strong> с <ins>рич-форматированием</ins></p>",
		},
		{
			Name:     "25",
			State:    `{"blocks":[{"type":"unstyled","inlineStyleRanges":[],"entityRanges":[],"data":{},"text":"Привет, Life!","depth":0}],"entityMap":{}}`,
			Expected: "<p>Привет, Life!</p>",
		},
		{
			Name:     "26",
			State:    `{"blocks":[{"key":"16tag","text":"qwertyuiopasdfghjkl","type":"unstyled","depth":0,"inlineStyleRanges":[{"offset":0,"length":19,"style":"CODE"},{"offset":0,"length":12,"style":"BOLD"},{"offset":1,"length":14,"style":"STRIKETHROUGH"},{"offset":4,"length":12,"style":"ITALIC"}],"entityRanges":[{"offset":3,"length":15,"key":0}],"data":{}}],"entityMap":{"0":{"type":"LINK","mutability":"MUTABLE","data":{"url":"ya.ru"}}}}`,
			Expected: `<p><code><strong>q</strong></code><code><strong><del>we</del></strong></code><a href="ya.ru" target="_blank"><code><strong><del>r</del></strong></code></a><a href="ya.ru" target="_blank"><code><strong><del><em>tyuiopas</em></del></strong></code></a><a href="ya.ru" target="_blank"><code><del><em>dfg</em></del></code></a><a href="ya.ru" target="_blank"><code><em>h</em></code></a><a href="ya.ru" target="_blank"><code>jk</code></a><code>l</code></p>`,
		},
		{
			Name:     "One symbol",
			State:    `{"entityMap":{"0":{"type":"LINK","data":{"url":"http://example.com"}}},"blocks":[{"text":"Q","type":"unstyled","depth":0,"inlineStyleRanges":[{"offset":0,"length":1,"style":"BOLD"}],"entityRanges":[{"offset":0,"length":1,"key":0}]}]}`,
			Expected: `<p><a href="http://example.com" target="_blank"><strong>Q</strong></a></p>`,
		},
		{
			Name:     "Wrong Ranges",
			State:    `{"entityMap":{"0":{"type":"LINK","data":{"url":"http://example.com"}}},"blocks":[{"text":"Q","type":"unstyled","depth":0,"inlineStyleRanges":[{"offset":5,"length":1,"style":"BOLD"}],"entityRanges":[{"offset":0,"length":1,"key":0}]}]}`,
			Expected: `<p><a href="http://example.com" target="_blank">Q</a></p>`,
		},
	}
}

func GetTestsPlainTable() []TestTable {
	return []TestTable{
		{
			Name:     "1",
			State:    `{"entityMap":{},"blocks":[{"key":"4g603","text":"dasdasdasdsadsaывфвыфв","type":"unstyled","depth":0,"inlineStyleRanges":[{"offset":0,"length":22,"style":"BOLD"}],"entityRanges":[],"data":{}}]}`,
			Expected: "dasdasdasdsadsaывфвыфв\n",
		},
		{
			Name:     "2",
			State:    `{"entityMap":{"0":{"type":"LINK","mutability":"MUTABLE","data":{"url":"ya.ru"}}},"blocks":[{"key":"4g603","text":"dasdasdasdsadsaывфвыфв","type":"unstyled","depth":0,"inlineStyleRanges":[],"entityRanges":[{"offset":12,"length":2,"key":0}],"data":{}}]}`,
			Expected: "dasdasdasdsadsaывфвыфв\n",
		},
		{
			Name:     "3",
			State:    `{"entityMap":{},"blocks":[{"key":"18u09","text":"\"Война и мир\" на экранах Би-Би-Си","type":"unstyled","depth":0,"inlineStyleRanges":[],"entityRanges":[],"data":{}}]}`,
			Expected: "\"Война и мир\" на экранах Би-Би-Си\n",
		},
	}
}

const ExampleDraftStateSource = `{
	  "entityMap": {
	    "0": {
	      "type": "LINK",
	      "mutability": "MUTABLE",
	      "data": {
		"url": "https://medium.com/@rajaraodv/how-draft-js-represents-rich-text-data-eeabb5f25cf2#.ce9y2wyux"
	     }}},
	  "blocks": [{
	      "text": "Rich text with link",
	      "type": "unstyled",
	      "depth": 0,
	      "inlineStyleRanges": [
		{
		  "offset": 0,
		  "length": 4,
		  "style": "BOLD"
		},
		{
		  "offset": 2,
		  "length": 10,
		  "style": "UNDERLINE"
		},
		{
		  "offset": 5,
		  "length": 4,
		  "style": "ITALIC"
		},
		{
		  "offset": 10,
		  "length": 4,
		  "style": "CODE"
		}
	      ],
	      "entityRanges": [{
		  "offset": 15,
		  "length": 4,
		  "key": 0
		}],
	      "data": {}
	    }]}`
