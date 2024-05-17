"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.PdfTable = void 0;
class PdfTable {
    constructor(obj) {
        this.tableNumber = obj.tableNumber;
        this.numrows = obj.numrows;
        this.numcols = obj.numcols;
        this.data = obj.data;
    }
    asDelimitedText(separator = ',') {
        let out = '';
        this.data.forEach(line => {
            const text = line.map(str => {
                if (str.includes('"'))
                    return `"${str.replace(/"/g, '""')}"`;
                else if (str.match(/[\s,]/))
                    return `"${str}"`;
                else
                    return str;
            }).join(separator);
            writeLn(text);
        });
        return out;
        function writeLn(text = '') {
            out += text + '\r\n';
        }
    }
    asHtml() {
        let out = '<table>';
        const data = this.data;
        for (let i = 0; i < this.numrows; i++) {
            let text = '<tr>';
            for (let j = 0; j < this.numcols; j++)
                text += '<td>' + (data[i][j] || '&nbsp;').replace(/</g, '&lt;') + '</td>';
            text += '</tr>';
            out += text;
        }
        out += '</table>';
        return out;
    }
}
exports.PdfTable = PdfTable;
//# sourceMappingURL=PdfTable.js.map