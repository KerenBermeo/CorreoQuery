export interface SearchResult {
    _id: string;
    _source: {
      subject: string;
      from: string;
      to: string[];
      date: string;
      content: string;
    };
}