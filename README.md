## Вариант задания
Реализация — `antlr`. Целевой код — `go`(`gccgo`).

### Язык
Графовый язык.
Встроенные типы — `node`, `arc`, `graph`.
Переопределить операции — `+`, `-`, `*`, `\`, и т.д. для встроенных типов.

#### Дополнительные свойства языка

_(2)_ Неявное объявление переменных  
_(1)_ Явное преобразование `a=(int)b`  
_(1)_ Одноцелевое присваивание `a=b`  
_(1)_ Структуры, ограничивающие область видимости — подпрограммы  
_(2)_ Неявный маркер блока(как в python)  
_(1)_ Двух вариантный оператор условий `if-then-else`  
_(1)_ Перегрузка подпрограмм отсутствует  
_(1)_ Передача параметров в подпрограмму только по значению и возращаемому значению  
_(1)_ Объявление подпрограмм только в начале файла  

Дополнительные размышления:
1. Взвешенные дуги
1. Приведение только узла в вершину и, может быть, графа из одной вершины в вершину
1. В вешину только её название и, может быть, вес(а может и содержимое будет и операции над ними)
1. Стандартные функции с операциями над графами
1. Писать без main

Операции над графами:
1. Тензорное произведение
1. Проверка графа на цикличность