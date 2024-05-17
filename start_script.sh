echo 'Running script'
cd node-pdf-table-parser
yarn install
# ponnoo
yarn start company-details.pdf
# vindhya
# yarn start Vindhya-company.pdf
cd ..
cd go-sheets-writer
go run main.go tableData.json
cd ..
