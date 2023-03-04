# go-test-task

### Ферма

На ферме живут курицы.
Известно, что курица несет от 2 до 5 яиц каждые 2-10 секунд (курицы несут яйца независимо друг от друга).
Все снесенные яйца складываются в холодильник. И, через 10-15 секунд приходит фермер за определенным количеством яиц (случайное число от 10-20). Если в холодильнике достаточное количество яиц, он забирает необходимое, иначе забирает все яйца из холодильника, если их было меньше.

##### Входные параметры
- Диапазоны (количества яиц и временных промежутков) задаются при старте;
- Число из диапазонов выбирается случайно во время выполнения программы.

##### Задание
Необходимо реализовать программу эмулирующую алгоритм взаимодействия, описанный в условии. Дополнительно сделать http-handler для получения текущего количества яиц в холодильнике.

- Ожидается, что вместе с проектом будет приложен README с описанием как запустить и проверить работоспособность приложения;
- Расположить код в виде оформленного github-репозитория;
- *Примерное* время выполнения задания: 2 часа.

###### Преимуществом будет
- Разумное использование сторонних библиотек;
- Логирование;
- Unit-тесты;
- Все константы, настраиваемые через конфиг;
- Оформленный Makefile;
- Сборка Docker-образа с приложением;
- Логическое разделение на коммиты.