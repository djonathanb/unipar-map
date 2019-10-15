# unipar-map
Minimum Path API of Environments at UNIPAR

## Running Locally

`go run src/main.go`

## API Resources

> Access the environment [here](https://arcane-ravine-32235.herokuapp.com/environments/)

### List environments

`/environments/`

### Route between environments

`/from/{env}/to/{env}`

**Example**: [/from/CANTINA/to/TECNICA/](https://arcane-ravine-32235.herokuapp.com/from/CANTINA/to/TECNICA/)

### Route to nearest bathroom

`/bathroom/{from}/`

**Example**: [/bathroom/FIES/](https://arcane-ravine-32235.herokuapp.com/bathroom/FIES/)
