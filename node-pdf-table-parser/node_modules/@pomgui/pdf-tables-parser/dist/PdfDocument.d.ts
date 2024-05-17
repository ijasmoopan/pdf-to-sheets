/// <reference types="node" />
import { Options, PdfPage } from './types';
export declare class PdfDocument {
    private _options;
    numPages: number;
    pages: PdfPage[];
    constructor(_options?: Options);
    load(source: string | Buffer): Promise<void>;
    private _extractTables;
    private _splitTables;
    private _extractRows;
    private _extractNextRow;
}
