Проект содержит набор скриптов и утилит для воспроизведения музыки в Sky: Children of the Light

# Требования

- Android устройство
- Включенный режим разработчика
- Linux и инструменты adb

# Флаги и запуск

## Обязательные
- **serial** - ADB ID вашего устройства.
    Можно узнать выполнив `adb devices`
- **track** - Путь к композиции

## Опциональные
- **adb** - Путь расположения adb утилиты.
- **speed** - Общий множитель для всех задержек.
    Например, 1.5 ускорит композицию в полтора раза.
- **t** - Множитель для таймингов. По умолчанию 200мс.
- **shift** - Сдвиг нот, для сдвига на октаву укажите 7.

Пример для запуска композиции:
```
./player --serial 6a58b007 --track ./tracks/test.txt
```

Пример для настройки таймингов, увелечения скорости и сдвига по ноте:
```
./player --serial 6a58b007 --track ./tracks/test.txt --timing 300 --speed 1.5 --shift 1
```

# Представление композиции

Есть 3 вида блоков в композиции:
- **Задержка** - может представляться в нескольких вариантах:
    - **Число** - будет воспринято как миллисекунды: `200 500`
    - **tN** - будет воспринято как 200 * N: `t t5 t10`.
        Множитель можно изменить флагом `-t`
    - **Тире** - несколько последовательных тире,
        будет воспринято как (200 * количество): `- -- ---`.
        Множитель можно изменить флагом `-t`
- **Нота** - буквенно-цифровое представление ноты: `C4, A2 G5`
- **Аккорд** - буквенное представление аккорда: `Am, E, G`

В общем случае для задержки рекомендуется использовать `timing`
или тире. В таком случае вы сможете менять время такта.

Кроме этого поддерживаются комментарии, для этого следует начать
строку с символа `#`.

И управляющие комментарии, позволяющие настроить стандартный
тайминг и сдвиг для композиции.

```
#!TIMING:300
#!SHIFT:2
```

# Транспонирование тона

Утилита сама транспонирует октавы.
Например, если вы используете ноты `A4 B5 C4`, то все ноты на 4 октаве
будут транспонированы до первой октавы в игре, а 5 октава станет 2.
```
[C4] [  ] [  ] [  ] [  ]
[A4] [  ] {  } {  } {  }
{  } {  } {  } {B5} (  )
```

Если вы используете только ноты второй октавы и хотите играть
исключительно на них, тогда используйте флаг `shift` со значением `7`

Например, без флага нота `A4` будет играть так:
```
[  ] [  ] [  ] [  ] [  ]
[A4] [  ] {  } {  } {  }
{  } {  } {  } {  } (  )
```
Со сдвигом октавы на вторую так:
```
[  ] [  ] [  ] [  ] [  ]
[  ] [  ] {  } {  } {  }
{  } {  } {A4} {  } (  )
```

## Примеры композиций

Можно найти в директории tracks