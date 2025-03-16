// src/utils/date-formatter.js (or .ts)
export function formatDate(date: string | number | Date): string {
    // Your date formatting logic here
    return new Date(date).toLocaleDateString();
  }  