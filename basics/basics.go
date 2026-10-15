package basics

// Wird ein importiertes Package nicht genutzt, wirft der Compiler einen Fehler. Prinzipiell ist die Philosophie von Go, gezwungenermaßen sauberen Code zu schreiben.
import (
	"errors"
	"fmt"
	"sync"
)

func Variablen() {

	// Der Typ einer Variable muss für den Compiler sofort bei der Deklaration erkennbar sein.
	// In Go wird prinzipiell aber eigentlich nichts deklariert, sondern immer auch sofort initialisiert.
	var test1 int = 13 // Vollständige Schreibweise
	var test2 = 14     // Zuweisung impliziert Datentyp
	test3 := 13        // Kurzschreibweise für var test1 int = 13
	var test4 int      // "Deklaration". Der Compiler weist im Hintergrund leeren Variablen immer einen Defaultwert zu. 0 z.B. bei int (ist je nach Plattform 32 oder 64 Bit groß), uint8,

	// Das Nutzen von _ = varname signalisiert dem Compiler, das die Variable ungenutzt existieren darf (sollte temporär zur Entwicklung genutzt werden und am ende entfernt werden,
	// damit der Code am Ende sauber ist und keine überflüssigen Variablen mehr vorhanden sind).
	_ = test1
	_ = test2
	_ = test3

	fmt.Println(test4) // Default Zero-Value für INT = 0

	var s string = "Hello World!" // Zero-Value für INT = ""
	var b uint8 = 0b00001111      // Zero-Value für uint8 = 0
	var b2 byte = 0b0000_1111     // Zero-Value für byte = 0, byte = uint8. 0000_1111 lesbare erlaubte Schreibweise für Compiler
	var f32 float32 = 3.1415      // Zero-Value für float32 = 0
	var f64 float64 = 3.1415      // Zero-Value für float64 = 0
	var ok bool = true            // Zero-Value für bool = false

	fmt.Println(s)
	fmt.Println(f32)
	fmt.Println(f64)
	fmt.Println(ok)
	_ = b2

	// Binärzahlen printen
	fmt.Println(b)
	fmt.Printf("%08b\n", b)
	fmt.Printf("%02X\n", b)
	fmt.Printf("%#08b\n", b)

}

func Kontrollstrukturen() {

	for i := 0; i < 3; i++ {
		fmt.Println("Takt: ", i)
	}

	// while Ersatz
	var n int = 0
	for n < 3 {
		fmt.Println(n)
		n++
	}

	// If mit initialisierungsstatement, also zu vergleichenden Wert im If-Statement
	if length := len("Test"); length >= 4 {
		fmt.Println(length)
	}

	// Switch - Kein break nötig, passiert automatisch.
	state := "RUNNING"

	switch state {
	case "STOPPED":
		fmt.Println("Anhalten")
	case "RUNNING", "IDLE":
		fmt.Println("Ready")
	default:
		fmt.Println("Unbekannter Status")
	}

}

// Fehlerbehandlung
// Wird in Go über direkte Prüfungen innerhalb der Funktion/aufrufenden Funktion ebenenweise durchgeführt. Es gibt kein Try/Catch
// Bei einer Funktion werden immer zwei Werte returned. Einmal der tatsächliche Rückgabewert und einmal nil (kein Error), oder eine Fehlermeldung.
// Es wird in der Aufrufenden Funktion dann in jeder Schichte geprüft, ob der zweite Rückgabewert = oder != nil ist.
func leseRegister(adresse int) (float64, error) {

	if adresse == -1 {
		return 0, errors.New("ungültige Adresse")
	}

	return 42, nil
}

func LeseTemperatur() (float64, error) {

	raw, err := leseRegister(-1)

	if err != nil {
		return 0, fmt.Errorf("Temperatur lesen: %w", err)
	}
	return float64(raw) / 10, nil

}

func Arrays() {

	var a [3]int // => 0,0,0
	b := [3]int{1, 2, 3}
	c := [...]int{1, 2, 3, 4} // Compiler zählt die Größe selbst

	a[0] = 10
	fmt.Println(a, b, c)
}

func Slices() {
	// Slices sind Go's dynamische Arrays
	s := []int{1, 2, 3}
	s = append(s, 4, 5)

	fmt.Println(s, len(s), cap(s)) // cap zählt die Anzahl an reserviertem Speicher. Kann auf Arrays oder Slices angewendet werden. Man kann beim Hinzufügen von Elementen als sehen, wie ein Slice neuen Speicher im Hintergrund alloziert.
	fmt.Println(s[1:])             // List Slicing, wie in Python

	for i, v := range s {
		fmt.Println(i, v)
	}

	// Wenn der Iterator nicht gebraucht wird, muss er durch _ ersetzt werden, sonst meckert der Compiler.
	for _, v := range s {
		fmt.Println(v)
	}

}
func Maps() {

	status := map[string]int{
		"Presse": 1,
		"Säge":   2,
	}

	status["Fräse"] = 1 // Überschreiben, hinzufügen
	delete(status, "Fräse")
	fmt.Println(len(status))

	// ok ist eine Variable, die innerhalb eines if-Statements existiert.
	// Das ok wird durch Abfragen von status["Bohrer"] gesetzt. Das zweite ok ist praktisch die Prüfung. Es wird also quasi das geprüft: if true oder if false.

	if wert, ok := status["Bohrer"]; ok {
		fmt.Println("gefunden:", wert, ok)
	} else {
		fmt.Println("nicht vorhanden", ok)
	}

	for name, s := range status {
		fmt.Println(name, s)
	}
}

type Machine struct {
	Name       string
	Temperatur float64
	Aktiv      bool
	Tags       []string
}

type MachinePark struct {
	Location string
	Machines []Machine
}

func Structs() {

	m := Machine{
		Name:       "Presse",
		Temperatur: 42.5,
		Aktiv:      false,
	}

	m.Temperatur = 45.0

	fmt.Printf("%+v\n", m) // %v+ = Multiple values vermutlich

	p := &m // Pointer auf m
	p.Temperatur = 50.5

	kopie := m
	kopie.Name = "Kopie"
	fmt.Println(m.Name, kopie.Name)

	m.Tags = []string{"T1", "T2", "T3", "T4", "T5"}
	fmt.Println(m.Tags)
}

func Kombinationen() {

	maschinen := []Machine{
		{Name: "Presse", Temperatur: 50},
		{Name: "Säge", Temperatur: 43.43},
	}

	maschinen = append(maschinen, Machine{Name: "Fräse"})

	for _, machine := range maschinen {
		machine.Aktiv = true
	}
	fmt.Println(maschinen) // machine ist nur eine Kopie

	for i := range maschinen {
		maschinen[i].Aktiv = true
	}
	fmt.Println(maschinen) // maschinen[i] ändert den tatsächlichen Wert

}

func GoRoutines() {

	// Parallelität und Nebenläufigkeit:
	// Ein Programm verfügt über mehrere Threads, die der OS-Scheduler auf die Kerne verteilt und sagt, wann die Threads des Prozesses an der Reihe sind.
	// Parallelität heißt, dass wenn ein Prozess aus z.B. 4 Threads besteht, der OS-Scheduler jeden Thread auf einen Kern verteilen würde, bei einer CPU mit mindestens vier Kernen.
	// Nebenläufigkeit heißt, das auf den jeweiligen Kernen, der OS-Scheduler Threads von Prozessen smart CPU-Zeit gibt, wodurch die Prozesse überlappen. Das bedeutet,
	// das z.B. bei zwei Prozessen (Go, Excel) die Threads nebenläufig, also z.B. im Wechsel verarbeitet werden. Auch wenn z.B. Go einen Prozess startet, der vier Threads hat,
	// kann es sein, dass alle vier Threads auf dem selben Kern verarbeiet werden, oder 3 Threads auf einem und einer auf einem Anderen, sprich sie würden dann nebenläufig verarbeitet werden.

	// Bei Python oder C kann zur Laufzeit des Programms ein weiterer Thread erzeugt werden. Wenn z.B. in Python ein Thread im Programm erzeugt wird,
	// handelt es sich um einen echten OS-Thread, den der OS-Scheduler verwaltet. Dieser Thread ist mehrere MB groß.
	// Prinzipiell wird jeder Prozess in Threads unterteilt vom Kernel. Beim Python Beispiel z.B. fünf Threads und zur Laufzeit würde ein sechster dann hinzugefügt werden.

	// In Go behält der Go-Prozess permanent seine fünf Threads. Das Pendant zu Pythons oder Cs Threads sind Go-Routines.
	// Wenn Go mehrere Go-Routines innerhalb des Programms startet, verwaltet der Go-Scheduler diese Go-Routines und weißt sie den Threads seines OS-Prozesses zu.
	// Der OS-Scheduler weißt die Threads eines OS-Prozesses also den Kernen zu und der Go-Scheduler weißt seine Go-Routines (analog zu OS-Threads) den OS-Threads (analog zu Kernen) zu.

	// go fmt.Println("Hallo aus der GoRoutine") // Wenn nur eine go-routine ausgeührt wird im Hauptprogramm und das Hauptprogramm nichts weiter tut, ist das Hauptprogramm schneller fertig,
	// als die go-routine, wodurch Diese dann nicht ausgeührt wird.

	// Es gibt drei Möglichkeiten, sicherzustellen, dass die go-routines vollständig durchlaufen
	// 1. in einer dauerhaften Anwendung, wie einem Webserver

	// 2. Wenn nur auf das Ende gewartet wird: sync.WaitGroup
	var wg sync.WaitGroup
	wg.Add(1) // Anzahl der go-routines eintragen (Zähler wird um 1 erhöht)

	go func() {
		defer wg.Done() // Melden, dass Funktion fertig ist (Zähler wird um 1 verringert)
		fmt.Println("Hallo aus der GoRoutine")
	}()
	wg.Wait() // Wait blockiert solange, bis der Wert = 0 ist. Also Add zählt hoch und Done zählt wieder runter und Wait prüft.

	// 3. Wenn das Ergebnis weiter verarbeitet werden muss innerhalb des Programms: Channel
	ch := make(chan string) // Ein Channel ist einfach praktisch einfach nur ein Rohr, durch das Daten fließen

	go func() {
		ch <- "Hallo vom anderen Ende des Rohrs"
	}()

	msg := <-ch // <-ch blockiert, bis ein Wert ankommt.
	fmt.Println(msg)

	// Channel kann auch mehrere Werte parken - Beispiel soll veranschaulichen, dass auch komplexe Berechnungen in eine Go-Routine ausgelagert werden können, sodass das Hauptprogramm weiter arbeiten kann.
	ch2 := make(chan int)

	go func() {
		for i := 1; i <= 3; i++ {
			ch2 <- i
		}
		close(ch2)
	}()

	for wert := range ch2 {
		fmt.Println("Empfangen aus CH2: ", wert)
	}

	// Mehrere Go-Routines schreiben in einen Channel und werden nebenläufig ausgeführt und durch den Go-Scheduler auf den OS-Threads des Hauptprozesses verteilt.
	ch3 := make(chan string)
	maschinen := []string{"Presse", "Säge", "Fräse"}

	for _, m := range maschinen {
		go func() {
			ch3 <- m + ": OK"
		}()
	}

	for range maschinen {
		fmt.Println(<-ch3) // Werden in irgendeiner Reihenfolge geprintet, je nachdem, welche Go-Routine schneller war.
		// ein Channel ist eine Queue (FiFo), also es wird in jeder Iteration das nächste Element konsumiert und der Channel wird kleiner.
	}

}
