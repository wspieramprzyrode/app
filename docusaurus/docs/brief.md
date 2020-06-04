---
id: brief
title: Opis systemu
---

## Załozenia ogolne

Aplikacja ma słuzyć  do  pomocy zarzadzania informacjami zwiazanymi z opieka nad roznymi aspektami przyrodniczymi.

Dostepna w formie ogolnodostpenj aplikacji mobilnej oraz strony www.

## Definicje

* Uzytkownik - kazda osoba korzystajaca z aplikacji mobilnej lub strony www bez koniecznosci posiadania na niej konta.
* zalogowany Uzytkownik - uzytkownik ktory dokonal rejestracji w aplikacji mobilnej lub na stronie www
* typ obiektu - zdefioniowana przez zarzadzajacymi serwisem
* obiekt - zdefioniowany w sytemie punkt posiadajace wspolrzedne geograficzne

## Funkcjonalności

### Zarzadzanie budkami lęgowymi ptaków

Aplikacja posiada mape na ktorej widoczna jest lokalizacja budek legowych ptakow.

Nie zalogowani uzytkownicy posiadają

#### Mapa budek

Mapa na ktorej znajduje sie lokalizacja budek legowych oparta na mapie openstreetmap.

Podczas otwarcia mapy pobierana jest z api lista budek w oparciu o aktualna lokalizacje uzytkownika. Domyslnie pokazywane mu sa budki znjadujace sie w okolicy.
W celu łatwiejszej identyfikacji budek poszczegolne typy budek posiadaja inne ikonki oraz kolory ikon w zaleznosci od tego czy dana budka posiada opiekuna lub nie i czy zostala jako uszkodzona.

#### Szczegoly budki

Po wybraniu budki legowej na mapie widza szczegoly dotyczace budki takie jak

* typ budki(o ile mozliwa jest jest jest klasyfikacja)
* ptaki jakie w nich sie najczesciej gniezdza - na podtswie typu budki
* informacje czy budka posiada opiekuna
* historie budki - czy zdjecia budki i jej stanu przeslane przez uzytkoniwkow
* przycisk umozliwiajacy zostanie opiekunen budki - wtedy potrzebne jest zalogowania lub zalozenie konta
* przycisk umozliwiajacy zgloszenie uszkodzenia budki - formularz z o oknam opisu i dolaczenia zdjec

### Zgłoszenia martwych lub chorych ptaków

Zalogowany uzytkownik posiada mozliwosc zglosznia napotkanych chorych lub martwych ptakow

### Dziennik oberwacji ptaków

Zalogowany uzytkownik moze stworzyc dziennik obserwacji ptakow podajac date i godziny , wskazujac miejsce obserwacji poprzez wskazanie go na mapie, gatunek zaoberwowanych ptakow. Moze dodac rowniez zdjecia.