package groupie

import (
    "encoding/json"
    "fmt"
    "io"
    "log"
    "net/http"
)

type Artist struct {
    ID           int      `json:"id"`
    Image        string   `json:"image"`
    Name         string   `json:"name"`
    Members      []string `json:"members"`
    CreationDate int      `json:"creationDate"`
    FirstAlbum   string   `json:"firstAlbum"`
    Locations    string   `json:"locations"`
    ConcertDates string   `json:"concertDates"`
    Relations    string   `json:"relations"`
}

func fetchArtists() ([]Artist, error) {
    const apiURL = "https://groupietrackers.herokuapp.com/api/artists"

    log.Printf("Fetching artists from: %s", apiURL)

    resp, err := http.Get(apiURL)
    if err != nil {
        return nil, fmt.Errorf("GET request failed: %w", err)
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
       return nil, fmt.Errorf("unexpected HTTP status: %d", resp.StatusCode)
    }

    body, err := io.ReadAll(resp.Body)
    if err != nil {
        return nil, fmt.Errorf("failed to read response body: %w", err)
    }

    var artists []Artist
    if err := json.Unmarshal(body, &artists); err != nil {
        return nil, fmt.Errorf("failed to decode JSON: %w", err)
    }

    log.Printf("Successfully fetched %d artists", len(artists))
    return artists, nil
}

func main() {
    artists, err := fetchArtists()
    if err != nil {
        log.Fatalf("Error: %v", err)
    }

    for _, artist := range artists {
        fmt.Printf("%d: %s\n", artist.ID, artist.Name)
    }
}
