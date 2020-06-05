---
id: brief
title: Opis systemu
---

## Założenia ogólne

Aplikacja ma służyć do pomocy zarządzania informacjami zwiazanymi z opieką nad rożnymi aspektami przyrodniczymi.

Dostępna w formie ogólnodostępnej aplikacji mobilnej oraz strony www.

## Definicje

### Role w systemie

#### użytkownik

każda osoba korzystająca z aplikacji mobilnej lub strony www bez konieczności posiadania w niej konta.

#### zalogowany użytkownik

[użytkownik](#użytkownik) który dokonał rejestracji w aplikacji mobilnej lub na stronie www oraz zweryfikował podane w procesie rejstracyjnym dane.

#### administrator systemu

[zalogowany użytkownik](#zalogowany-użytkownik) posiadający uprawnienia do wszystkich operacji z systemie.

#### opiekun obiektu

[zalogowany użytkownik](#zalogowany-użytkownik) który ma przypisane do siebie [obiekty](#obiekt) które ma pod swoim nadzorem

### Obiekty

#### typ obiektu

zdefiniowany przez zarządzającymi serwisem katalog rodzajów obiektów do której mogą zostać przypisane poszczególne obiekty np:

* budka lęgowa ptaków
* pomnik przyrody

#### obiekt

zdefiniowany w sytemie punkt posiadający współrzędne geograficzne oraz [typ obiektu](#typ-obiektu)

### Zwierząta

#### gatunek ptaka

zdefiniowany w słowniku ptaków zarządzanym przez [administratora serwisu](#administrator-systemu)

## Funkcjonalności

### Zarządzanie budkami lęgowymi ptaków

Aplikacja posiada mapę na której widoczna jest lokalizacja budek lęgowych ptaków.

#### Znakowanie budek

[Użytkownik](#użytkownik) posiada możliwość wygenerowania etykiet w postacie pdf zawierających qrcode zawierający unikatowy identyfikator obiektu.

Po wydrukowaniu użytkownik może oznaczyć obiekt przywieszając go w sposób umożliwiający korzystanie z niego z telefonów komórkowych. Użytkownik powinin uzyskać zgodę właściciela lub zarządzającego na temrenie którego znjaduje się obiekt.

Po zrobieniu fotografii tego kodu zostajemy przekierunkowani do aplikacji lub stronę www na cześć ze [szczegółowymi informacjami](#szczegóły-budki-lęgowej) dotyczącymi tego [obiektu](#obiekt).

W przypadku braku danych w systemie jako zalogowany użytkownik rozpoczynami tworzenie nowego obiektu. Użytkownik nie zalogowany proszony jest o stworzenie nowego konta lub zalogwanie się do systemu.

#### Mapa budek

Mapa na której znajduje się lokalizacja budek lęgowych oparta na openstreetmap. Dostępna jest ona dla każdego [użytkownika](#użytkownik) aplikacji.

Podczas otwarcia mapy pobierana jest z api lista budek w oparciu o aktualna lokalizację [użytkownika](#użytkownik). Domyślnie pokazywane mu są budki znajdujące sie w okolicy.
W celu łatwiejszej identyfikacji budek poszczególne typy budek posiadają inne ikonki oraz kolory ikon w zależności od tego czy dana budka posiada opiekuna lub nie i czy została zgłoszona jako uszkodzona.

#### Szczegóły budki lęgowej

Po wybraniu [budki lęgowej](#obiekt) na mapie [użytkownik](#użytkownik) widzi szczegóły dotyczące budki:

* typ budki (o ile możliwa jest jest jest klasyfikacja)
* ptaki jakie w nich się najcześciej pojawiają - na podstawie typu budki
* informacje czy [budka](#obiekt) posiada opiekuna
* historie [budki](#obiekt) - czyli zdjęcia [budki](#obiekt) i jej stanu przesłane przez użytkowników
* przycisk umożliwiający zostanie opiekunem [budki](#obiekt) - wtedy potrzebne jest zalogowania lub założenie konta
* przycisk umożliwiający [zgłoszenie uszkodzenia budki](#zgłoszenie-uszkodzenia-budki-lęgowej)
* przycisk umożliwiający [zgloszenie obserwacji ptaków wykorzytujących budkę](#informacja-o-zamieszkaniu-budki-l%C4%99gowej)
* przycisk umożliwiający [wygenerowanie etykiety z qrcode](#znakowanie-budek) dla danej budki lęgowej

#### Dodanie nowej budki lęgowej

[Zalogowany użytkownik](#zalogowany-użytkownik) może dodać nową budkę lęgową przez wypełnienie formularza:

* wybierając typ budki
* wskazanie na mapie miejsca z jak największą dokładnością
* dodanie zdjęcia budki w celu jej identyfikacji
* wybranie czy osoba zgłaszająca budkę jest jej opiekunem czy nie

W przypadku kiedy [zalogowany użytkownik](#zalogowany-użytkownik) wybiera opcję opiekun budki dodana budka zostaniej dodana do strony Moje budki. Zaczyna on również otrzymywać powiadomienia związene z ta budką.

#### Zgłoszenie uszkodzenia budki lęgowej

[Użytkownik](#użytkownik) posiada możliwość zgłoszenia uszkodzenia lub zniszcznie budki lęgowej poprzez podanie:

* wybierając uszkodzona lub zniszczona z listy wyboru
* daty
* opisu (opcjonalne)
* załączenie zdjecia

W przyapdku kiedy do danej budki lęgowej przypisany jest jej [opiekun](#opiekun-obiektu) informacja ta zostanie do niego wysłana w postaci notyfikacji push oraz wiadomości email. Budka takiej zostanie nadany status zniszczona lub uszkodzona w zależności od statusu.

W przypadku kiedy budka taka nie posiada opiekuna informacja taka zostaje wysłana do administratora systemu

#### Naprawa budki

[Opiekun budki](#opiekun-obiektu) lub [administrator serwisu](#administrator-systemu) może zmienić status danej budki lęgowej na normalną po naprawie budki załączajać  jej zdjęcie po pracach.

#### Usunięcie budki

[Opiekun budki](#opiekun-obiektu) lub [administrator serwisu](#administrator-systemu) może usunąć budkę lęgową podajac przyczynę usunięcia np wycinka drzew.

#### Rezygnacja z opieki nad budką

[Opiekun budki](#opiekun-obiektu) lub [administrator serwisu](#administrator-systemu) może zrezygnować z opieki nad budką

#### Zmiana opiekuna nad budką

[Opiekun budki](#opiekun-obiektu) lub [administrator serwisu](#administrator-systemu) może wysłać zaproszenia na adres email do przejęcia opieki nad budką lęgową.

Po wypisaniu formularza zostanie wysłany email na podany adres zawierający link do potwierdzenia przejęcia opieki.

Kiedy podany email jest podany w bazie użytkowników budka po potwierdzeniu przejęcia zmiania opiekuna.

W przypadku kiedy email nie jest zarejestowany - nowy potencjalny opiekun musi założyć konto w systemie i wtedy potwierdzić przejęcie opieki nad budką lęgową.

#### Moje budki

Lista budek lęgowych w których [zalogowany użytkownik](#zalogowany-użytkownik) jest przypisany jako [opiekun budki](#opiekun-obiektu).

#### Informacja o zamieszkaniu budki lęgowej

[Użytkownik](#użytkownik) może zgłosić obserwacje ptaków które zajęły budkę lęgową.
Na formularzu wpisuje:

* datę opbserwacji ptaków przy budce lęgowej
* gatunek ptaka
* posiada opcionalną możliwość dodania zdjęć

Informacja ta zostaje wysłana do opiekuna budki (o ile jest przypisany) oraz do administratorów systemu.

### Zgłoszenia martwych lub chorych ptaków

[Zalogowany użytkownik](#zalogowany-użytkownik) posiada możliwość zgłoszenia napotkanych chorych lub martwych ptaków poprzez formularz podając

* datę i godzinę
* wskazanie na mapie miejsca
* gatunek zaoberwowanych ptaków
* możliwości dołączenia zdjęć (opcjonalne)

### Dziennik oberwacji ptaków

#### Mapa obserwacji

[Użytkownik](#użytkownik) widzi na mapie obserwacje wszystkich [użytkowników](#użytkownik) którzy oznaczyli obserwacje jako publiczne.

#### Zarządzanie obserwacjami

[Zalogowany użytkownik](#zalogowany-użytkownik) może stworzyć dziennik obserwacji ptaków.

##### Dodanie obserwacji

Formularz zawierający:

* datę i godzinę obserwacji
* wskazanie na mapie miejsca obserwacji
* gatunek zaobserwowanych ptaków
* możliwości dołączenia zdjęć (opcjonalne)
* wybór czy obserwacja może być widoczna dla wszystkich [użytkowników](#użytkownik) czy jest ona prywatna (domyślnie publiczna)

#### Lista oberwacji

[Zalogowany użytkownik](#zalogowany-użytkownik) posiada dostęp do listy swoich oberwacji może wybrać jeden w celu zobaczenia szczegółów.

Wchodząc na szczegóły widzi podane dane. Posiada on możliwość wejścia w tryb edycji lub usunąć oberwację.

#### Edycja obserwacji

[Zalogowany użytkownik](#zalogowany-użytkownik) edytując swoją obserwacje może:

* dodawać lub usunąć zdjęcia
* zmienić widoczność oberwacji z publicznej na prywatną i odwrotnie
* zmienić gatunek ptaka
* zmienić datę i godzinę
* zmienić miejsce obserwacji poprzez wybór nowego miejsca z mapy

### Dzikie wysypiska śmieci

#### Zgłoszenie nielegalnego wysypiska śmieci

[Użytkownik](#użytkownik) posiada możliwość zgłoszenia nielegalnego wysypiska śmieci podając:

* daty
* wskazanie na mapie miejsca wysypiska
* opisu (opcjonalne)
* załączenie zdjęć

Informacja ta zostaje wysłana do administratora systemu w celu weryfikacji oraz poźniejszego wysłania do służb które odpowiadają na danym terenie.

### Pomniki przyrody

#### Mapa pominików przyrody

Na mapie [użytkownik](#użytkownik) widzi pomniki przyrody znajdujące się w jego okolicy na postawie aktualnej pozycji. Widoczne są tylko zweryfikowane przez administatora systemu pomniki przyrody

#### Dodanie pomnika przyrody

[Użytkownik](#użytkownik) posiada możliwość zgłoszenia pomnika przyrody:

* wskazanie na mapie miejsca
* opisu (opcjonalne)
* załączenie zdjęć

Informacje o dodaniu nowego pomnika przyrody wysyłana jest do administatora systemu w celu jej weryfikacji

#### Szczegóły pomnika przyrody

Po wybraniu na mapie pomnika przyrod widzi szczegóły:

* zdjęcia
* przycisk umożliwiający [zgłoszenie uszkodzenia pomnika](#zg%C5%82oszenie-informacji-o-zniszczeniu-lub-uszkodzeniu-pomnika-przyrody)

##### Zgłoszenie informacji o zniszczeniu lub uszkodzeniu pomnika przyrody

[Użytkownik](#użytkownik) może zgłosić uszkodzenie lub zniszczenie przyrody
Na formularzu wpisuje:

* datę
* opis
* posiada opcionalną możliwość dodania zdjęć

Informacja ta zostaje wysłana do administratorów systemu w celu weryfikacji i przekazania do odpowiednuch służb.
