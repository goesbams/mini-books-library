export interface Book {
  id: number;
  title: string;
  author: string;
  cover_image_url: string;
  description: string;
  publication_date: string;
  number_of_pages: number;
  isbn: string;
}

export interface CreateBookRequest {
  title: string;
  author: string;
  cover_image_url?: string;
  description?: string;
  publication_date: string;
  number_of_pages: number;
  isbn: string;
}

export interface UpdateBookRequest {
  title?: string;
  author?: string;
  cover_image_url?: string;
  description?: string;
  publication_date?: string;
  number_of_pages?: number;
  isbn?: string;
}

export interface ApiError {
  error: string;
  message: string | string[];
}

export interface ApiResponse<T> {
  data?: T;
  error?: ApiError;
}