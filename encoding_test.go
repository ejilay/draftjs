package draftjs

import (
	"encoding/json"
	"testing"
)

func TestDecorator_Render(t *testing.T) {
	contentStates := []ContentState{}
	var err error
	if err = json.Unmarshal([]byte(testString), &contentStates); err != nil {
		t.Errorf("Failed unmarshal content: %v", err)
		return
	}

	if err != nil {
		t.Error(err)
		return
	}

	config := DefaultConfig()
	i := 0
	for _, block := range contentStates {
		s := Render(&block, config)
		if s != needString[i] {
			t.Errorf("%s != %s", s, needString[i])
			return
		}
		i++
	}
}

var (
	testString string = `[
  {
    "entityMap": {

    },
    "blocks": [
      {
        "key": "4g603",
        "text": "dasdasdasdsadsaывфвыфв",
        "type": "unstyled",
        "depth": 0,
        "inlineStyleRanges": [
          {
            "offset": 0,
            "length": 22,
            "style": "BOLD"
          }
        ],
        "entityRanges": [

        ],
        "data": {

        }
      }
    ]
  },
  {
    "entityMap": {
      "0": {
        "type": "LINK",
        "mutability": "MUTABLE",
        "data": {
          "url": "ya.ru"
        }
      }
    },
    "blocks": [
      {
        "key": "4g603",
        "text": "dasdasdasdsadsaывфвыфв",
        "type": "unstyled",
        "depth": 0,
        "inlineStyleRanges": [

        ],
        "entityRanges": [
          {
            "offset": 12,
            "length": 2,
            "key": 0
          }
        ],
        "data": {

        }
      }
    ]
  },
  {
    "entityMap": {

    },
    "blocks": [
      {
        "key": "18u09",
        "text": "\"Война и мир\" на экранах Би-Би-Си",
        "type": "unstyled",
        "depth": 0,
        "inlineStyleRanges": [

        ],
        "entityRanges": [

        ],
        "data": {

        }
      }
    ]
  },
  {
    "entityMap": {

    },
    "blocks": [
      {
        "key": "ek0ec",
        "text": "Яндекс.Метро — интерактивная карта метро Москвы с расчётом времени и прокладкой маршрутов с учётом данных о закрытии станций и вес",
        "type": "unstyled",
        "depth": 0,
        "inlineStyleRanges": [
          {
            "offset": 7,
            "length": 5,
            "style": "ITALIC"
          },
          {
            "offset": 35,
            "length": 5,
            "style": "ITALIC"
          }
        ],
        "entityRanges": [

        ],
        "data": {

        }
      }
    ]
  },
  {
    "entityMap": {

    },
    "blocks": [
      {
        "key": "fdcqk",
        "text": "Тест 11 Тест 11 Тест 11  Тест 11  Тест 11  Тест 11  Тест 11 ",
        "type": "unstyled",
        "depth": 0,
        "inlineStyleRanges": [

        ],
        "entityRanges": [

        ],
        "data": {

        }
      }
    ]
  },
  {
    "entityMap": {

    },
    "blocks": [
      {
        "key": "ao1cv",
        "text": "Парам пам пам пам пам ",
        "type": "unstyled",
        "depth": 0,
        "inlineStyleRanges": [

        ],
        "entityRanges": [

        ],
        "data": {

        }
      }
    ]
  },
  {
    "blocks": [
      {
        "type": "unstyled",
        "inlineStyleRanges": [

        ],
        "entityRanges": [

        ],
        "data": {

        },
        "text": "Описание к изображению",
        "depth": 0
      }
    ],
    "entityMap": {

    }
  },
  {
    "blocks": [
      {
        "type": "unstyled",
        "inlineStyleRanges": [

        ],
        "entityRanges": [

        ],
        "data": {

        },
        "text": "Я прочитала книгу на английском, как только она появилась. На русском читать не буду, мне не нравится этот перевод.",
        "depth": 0
      }
    ],
    "entityMap": {

    }
  },
  {
    "blocks": [
      {
        "type": "unstyled",
        "inlineStyleRanges": [

        ],
        "entityRanges": [

        ],
        "data": {

        },
        "text": "Наталья Водянова, российская супермодель",
        "depth": 0
      }
    ],
    "entityMap": {

    }
  },
  {
    "blocks": [
      {
        "type": "unstyled",
        "inlineStyleRanges": [

        ],
        "entityRanges": [

        ],
        "data": {

        },
        "text": "Арам Ашотович Габрелянов",
        "depth": 0
      }
    ],
    "entityMap": {

    }
  },
  {
    "blocks": [
      {
        "type": "unstyled",
        "inlineStyleRanges": [

        ],
        "entityRanges": [

        ],
        "data": {

        },
        "text": "Не может быть идеального продукта. Если ты уверен, что создал идеальный продукт или идеальное СМИ, то тебе пора на пенсию.",
        "depth": 0
      }
    ],
    "entityMap": {

    }
  },
  {
    "blocks": [
      {
        "type": "unstyled",
        "inlineStyleRanges": [

        ],
        "entityRanges": [

        ],
        "data": {

        },
        "text": "",
        "depth": 0
      }
    ],
    "entityMap": {

    }
  },
  {
    "blocks": [
      {
        "type": "unstyled",
        "inlineStyleRanges": [

        ],
        "entityRanges": [

        ],
        "data": {

        },
        "text": "Текст карточки с картинкой",
        "depth": 0
      }
    ],
    "entityMap": {

    }
  },
  {
    "blocks": [
      {
        "type": "unstyled",
        "inlineStyleRanges": [

        ],
        "entityRanges": [

        ],
        "data": {

        },
        "text": "Абзац текста в карточке",
        "depth": 0
      }
    ],
    "entityMap": {

    }
  },
  {
    "blocks": [
      {
        "type": "unstyled",
        "inlineStyleRanges": [

        ],
        "entityRanges": [

        ],
        "data": {

        },
        "text": "Заголовок карточки",
        "depth": 0
      }
    ],
    "entityMap": {

    }
  },
  {
    "blocks": [
      {
        "type": "unstyled",
        "inlineStyleRanges": [

        ],
        "entityRanges": [

        ],
        "data": {

        },
        "text": "Арам Ашотович Габрелянов",
        "depth": 0
      }
    ],
    "entityMap": {

    }
  },
  {
    "blocks": [
      {
        "type": "unstyled",
        "inlineStyleRanges": [

        ],
        "entityRanges": [

        ],
        "data": {

        },
        "text": "Не может быть идеального продукта. Если ты уверен, что создал идеальный продукт или идеальное СМИ, то тебе пора на пенсию.",
        "depth": 0
      }
    ],
    "entityMap": {

    }
  },
  {
    "blocks": [
      {
        "type": "ordered-list-item",
        "inlineStyleRanges": [

        ],
        "entityRanges": [

        ],
        "data": {

        },
        "text": "1 элемент нумерованного списка",
        "depth": 0
      },
      {
        "type": "ordered-list-item",
        "inlineStyleRanges": [

        ],
        "entityRanges": [

        ],
        "data": {

        },
        "text": "2 элемент нумерованного списка",
        "depth": 0
      }
    ],
    "entityMap": {

    }
  },
  {
    "blocks": [
      {
        "type": "unordered-list-item",
        "inlineStyleRanges": [

        ],
        "entityRanges": [

        ],
        "data": {

        },
        "text": "1 элемент ненумерованного списка",
        "depth": 0
      },
      {
        "type": "unordered-list-item",
        "inlineStyleRanges": [

        ],
        "entityRanges": [

        ],
        "data": {

        },
        "text": "2 элемент ненумерованного списка",
        "depth": 0
      }
    ],
    "entityMap": {

    }
  },
  {
    "blocks": [
      {
        "type": "unstyled",
        "inlineStyleRanges": [
          {
            "length": 5,
            "offset": 0,
            "style": "ITALIC"
          },
          {
            "length": 6,
            "offset": 6,
            "style": "BOLD"
          },
          {
            "length": 19,
            "offset": 15,
            "style": "UNDERLINE"
          }
        ],
        "entityRanges": [

        ],
        "data": {

        },
        "text": "Абзац текста с рич-форматированием",
        "depth": 0
      }
    ],
    "entityMap": {

    }
  },
  {
    "blocks": [
      {
        "type": "unstyled",
        "inlineStyleRanges": [

        ],
        "entityRanges": [

        ],
        "data": {

        },
        "text": "Привет, Life!",
        "depth": 0
      }
    ],
    "entityMap": {

    }
  },
  {
    "blocks": [
      {
        "key": "16tag",
        "text": "qwertyuiopasdfghjkl",
        "type": "unstyled",
        "depth": 0,
        "inlineStyleRanges": [
          {
            "offset": 0,
            "length": 19,
            "style": "CODE"
          },
          {
            "offset": 0,
            "length": 12,
            "style": "BOLD"
          },
          {
            "offset": 1,
            "length": 14,
            "style": "STRIKETHROUGH"
          },
          {
            "offset": 4,
            "length": 12,
            "style": "ITALIC"
          }
        ],
        "entityRanges": [
          {
            "offset": 3,
            "length": 15,
            "key": 0
          }
        ],
        "data": {

        }
      }
    ],
    "entityMap": {
      "0": {
        "type": "LINK",
        "mutability": "MUTABLE",
        "data": {
          "url": "ya.ru"
        }
      }
    }
  },
  {
    "entityMap": {

    },
    "blocks": [
      {
        "key": "8lobd",
        "text": "a",
        "type": "unordered-list-item",
        "depth": 0,
        "inlineStyleRanges": [

        ],
        "entityRanges": [

        ],
        "data": {

        }
      },
      {
        "key": "7pppu",
        "text": "b",
        "type": "unordered-list-item",
        "depth": 0,
        "inlineStyleRanges": [

        ],
        "entityRanges": [

        ],
        "data": {

        }
      },
      {
        "key": "fn3d3",
        "text": "c",
        "type": "unordered-list-item",
        "depth": 0,
        "inlineStyleRanges": [

        ],
        "entityRanges": [

        ],
        "data": {

        }
      },
      {
        "key": "50oms",
        "text": "1",
        "type": "ordered-list-item",
        "depth": 1,
        "inlineStyleRanges": [

        ],
        "entityRanges": [

        ],
        "data": {

        }
      },
      {
        "key": "3rh3g",
        "text": "2",
        "type": "ordered-list-item",
        "depth": 1,
        "inlineStyleRanges": [

        ],
        "entityRanges": [

        ],
        "data": {

        }
      },
      {
        "key": "7mqrj",
        "text": "3",
        "type": "ordered-list-item",
        "depth": 1,
        "inlineStyleRanges": [

        ],
        "entityRanges": [

        ],
        "data": {

        }
      }
    ]
  }
]`

	needString = []string{
		`<p><strong>dasdasdasdsadsaывфвыфв</strong></p>`,
		`<p>dasdasdasdsa<a href="ya.ru" target="_blank">ds</a>aывфвыфв</p>`,
		`<p>&#34;Война и мир&#34; на экранах Би-Би-Си</p>`,
		`<p>Яндекс.<em>Метро</em> — интерактивная карта <em>метро</em> Москвы с расчётом времени и прокладкой маршрутов с учётом данных о закрытии станций и вес</p>`,
		`<p>Тест 11 Тест 11 Тест 11  Тест 11  Тест 11  Тест 11  Тест 11 </p>`,
		`<p>Парам пам пам пам пам </p>`,
		`<p>Описание к изображению</p>`,
		`<p>Я прочитала книгу на английском, как только она появилась. На русском читать не буду, мне не нравится этот перевод.</p>`,
		`<p>Наталья Водянова, российская супермодель</p>`,
		`<p>Арам Ашотович Габрелянов</p>`,
		`<p>Не может быть идеального продукта. Если ты уверен, что создал идеальный продукт или идеальное СМИ, то тебе пора на пенсию.</p>`,
		`<p></p>`,
		`<p>Текст карточки с картинкой</p>`,
		`<p>Абзац текста в карточке</p>`,
		`<p>Заголовок карточки</p>`,
		`<p>Арам Ашотович Габрелянов</p>`,
		`<p>Не может быть идеального продукта. Если ты уверен, что создал идеальный продукт или идеальное СМИ, то тебе пора на пенсию.</p>`,
		`<ol><li>1 элемент нумерованного списка</li><li>2 элемент нумерованного списка</li></ol>`,
		`<ul><li>1 элемент ненумерованного списка</li><li>2 элемент ненумерованного списка</li></ul>`,
		`<p><em>Абзац</em> <strong>текста</strong> с <ins>рич-форматированием</ins></p>`,
		`<p>Привет, Life!</p>`,
		`<p><code><strong>q</strong></code><code><strong><del>we</del></strong></code><a href="ya.ru" target="_blank"><code><strong><del>r</del></strong></code></a><a href="ya.ru" target="_blank"><code><strong><del><em>tyuiopas</em></del></strong></code></a><a href="ya.ru" target="_blank"><code><del><em>dfg</em></del></code></a><a href="ya.ru" target="_blank"><code><em>h</em></code></a><a href="ya.ru" target="_blank"><code>jk</code></a><code>l</code></p>`,
		`<ul><li>a</li><li>b</li><li>c<ol><li>1</li><li>2</li><li>3</li></ol></li></ul>`,
	}
)
