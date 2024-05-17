import { IPdfTable } from './types';
export declare class PdfTable implements IPdfTable {
    tableNumber: number;
    numrows: number;
    numcols: number;
    data: string[][];
    constructor(obj: IPdfTable);
    asDelimitedText(separator?: string): string;
    asHtml(): string;
}
