export interface Email {
    id: string;
    from: string; 
    to: string[]; 
    subject: string; 
    date: string; 
    content: string
}