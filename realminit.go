//This software converts a realm database exported csv file into javascript code that preloads the data.

package main

import "encoding/csv"
import "os"
import "fmt"

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
	Realm.open({schema: [sightingSchema, treeSchema]})

    .then(realm => {
`))
		
		
		var record []string
		var first bool = true
		for record, err = reader.Read(); err == nil; record, err = reader.Read() {
				
				if first {
					first = false
					continue
				}
			
				output.Write([]byte(`
      realm.write(() => {

        const CabbageTree = realm.create('TreeInfo', {

          id: `+record[0]+`,

          commonName:  '`+record[1]+`',

          maoriName:  '`+record[2]+`',

          latinName:  '`+record[3]+`',

          synonyms:  '`+record[4]+`',

          thumbnail:  '`+record[5]+`',

          image:  '`+record[6]+`',

          image2:  '`+record[7]+`',

          favourites:   `+record[8]+`,

          favouritesColour:  '`+record[9]+`',

          family: '`+record[10]+`',

          group: '`+record[11]+`',

          medicinal: `+record[12]+`,

          medicinalInfo: '`+record[13]+`',

          fruiting: '`+record[14]+`',

          poisonous: `+record[15]+`,

          poisonousInfo: '`+record[16]+`',

          flowering: '`+record[17]+`',

          speciesFeatures: '`+record[18]+`',

          description: '`+record[19]+`',

          didYouKnow: '`+record[20]+`',

          distribution: '`+record[21]+`',

          etymology: '`+record[22]+`',

          leafEdges: '`+record[23]+`',

          leafType: '`+record[24]+`',

          leafMarginSpecialty: '`+record[25]+`',

          leafSurfaceSpecialty: '`+record[26]+`',

          leafArrangement: '`+record[27]+`',

          leafShape: '`+record[28]+`',

          leafTip: '`+record[29]+`',

          flowerColour: '`+record[30]+`',

          fruitColour: '`+record[31]+`',

          fruitType: '`+record[32]+`',

          trunk: '`+record[33]+`',

          trunkColor: '`+record[34]+`',

        });

      });

				`))
		}
		
		output.Write([]byte(`
    });
		`))
		
		if fmt.Sprint(err) != "EOF" {
			fmt.Print("[ERROR] ")
			fmt.Println(err)
		}
	}
}
