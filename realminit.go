//This software converts a realm database exported csv file into javascript code that preloads the data.

package main

import "encoding/csv"
import "os"
import "fmt"
import "strings"

func main() {
	if len(os.Args) <= 1 {
		fmt.Println("Usage: realminit input.csv output.js")
	}

	if len(os.Args) > 1 {
		var output_name = "output.js"
		if len(os.Args) > 2 {
			output_name = os.Args[2]
		}

		fmt.Println(os.Args[1])

		file, err := os.Open(os.Args[1])
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		reader := csv.NewReader(file)
		output, err := os.Create(output_name)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

	output.Write([]byte(`
		const sightingSchema = {
		  name: 'Sighting',
		  properties: {
		    treeName:  'string',
		    location: 'string',
		    address: 'string',
		    picture: 'string',
		    notes: 'string',
		    date: 'string',
		     key: {type: 'int', default: 0},
		  }
		};

		const treeSchema = {
			name: 'TreeInfo',
		    primaryKey: 'id',
		    properties: {
		    
		    id: 'int',
		    commonName:  'string',
		    maoriName:  'string',
		    latinName:  'string',
		    synonyms:  'string',
		    thumbnail:  'string',
		    image2:  'string',
		    image3:  'string',
		    image4:  'string',
			image5:  'string',
			image6:  'string',
			
		    favourites:  {type: 'int', default: 0},

		    favouritesColour:  'string',
		    family: 'string',
		    group: 'string',
		    medicinal: {type: 'int', default: 0},
		    medicinalInfo: 'string',
		    fruiting: 'string',
		    poisonous: {type: 'int', default: 0},
		    poisonousInfo: 'string',
		    flowering: 'string',
		    speciesFeatures: 'string',
		    description: 'string',
		    didYouKnow: 'string',
		    distribution: 'string',
		    etymology: 'string',
		    leafEdges: 'string',
		    leafType: 'string',
		    leafMarginSpecialty: 'string',
		    leafSurfaceSpecialty: 'string',
		    leafArrangement: 'string',
		    leafShape: 'string',
		    leafTip: 'string',
		    flowerColour: 'string',
		    fruitColour: 'string',
		    fruitType: 'string',
		    trunk: 'string',
		    trunkColor: 'string',
		    sightings: {type: 'list', objectType: 'Sighting'},
		  }
		};

		export function populate() {
	Realm.open({schema: [sightingSchema, treeSchema]})

    .then(realm => {

			      realm.write(() => {
`))


		var record []string
		var first bool = true
		for record, err = reader.Read(); err == nil; record, err = reader.Read() {

				if first {
					first = false
					continue
				}

				for i, r := range record {
					r = strings.Replace(r, "\n","\\n", -1);
					r = strings.Replace(r, "\r","\\r", -1);
					record[i] = strings.Replace(r, "'","\\'", -1);
				}

				output.Write([]byte(`
        realm.create('TreeInfo', {

          id: `+record[0]+`,

          commonName:  '`+record[1]+`',

          maoriName:  '`+record[2]+`',

          latinName:  '`+record[3]+`',

          synonyms:  '`+record[4]+`',

          thumbnail:  '`+record[5]+`',

          image2:  '`+record[6]+`',
		  
		  image3:  '`+record[7]+`',

          image4:  '`+record[8]+`',
		  
		  image5:  '`+record[9]+`',

          image6:  '`+record[10]+`',

          favourites:   `+record[11]+`,

          favouritesColour:  '`+record[12]+`',

          family: '`+record[13]+`',

          group: '`+record[14]+`',

          medicinal: `+record[15]+`,

          medicinalInfo: '`+record[16]+`',

          fruiting: '`+record[17]+`',

          poisonous: `+record[18]+`,

          poisonousInfo: '`+record[19]+`',

          flowering: '`+record[20]+`',

          speciesFeatures: '`+record[21]+`',

          description: '`+record[22]+`',

          didYouKnow: '`+record[23]+`',

          distribution: '`+record[24]+`',

          etymology: '`+record[25]+`',

          leafEdges: '`+record[26]+`',

          leafType: '`+record[27]+`',

          leafMarginSpecialty: '`+record[28]+`',

          leafSurfaceSpecialty: '`+record[29]+`',

          leafArrangement: '`+record[30]+`',

          leafShape: '`+record[31]+`',

          leafTip: '`+record[32]+`',

          flowerColour: '`+record[33]+`',

          fruitColour: '`+record[34]+`',

          fruitType: '`+record[35]+`',

          trunk: '`+record[36]+`',

          trunkColor: '`+record[37]+`',

        });

				`))
		}

		output.Write([]byte(`
    });
		});
	}
		`))

		if fmt.Sprint(err) != "EOF" {
			fmt.Print("[ERROR] ")
			fmt.Println(err)
		}
	}
}
