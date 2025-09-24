import axios from 'axios';
import { Book, CreateBookRequest, UpdateBookRequest } from '@/types/book';

const API_BASE_URL = process.env.NEXT_PUBLIC_API_URL || 'http://localhost:9000';

const api = axios.create({
  baseURL: API_BASE_URL,
  headers: {
    'Content-Type': 'application/json',
  },
});

// Request interceptor for logging
api.interceptors.request.use(
  (config) => {
    console.log(`Making ${config.method?.toUpperCase()} request to ${config.url}`);
    return config;
  },
  (error) => {
    return Promise.reject(error);
  }
);

// Response interceptor for error handling
api.interceptors.response.use(
  (response) => {
    return response;
  },
  (error) => {
    console.error('API Error:', error.response?.data || error.message);
    return Promise.reject(error);
  }
);

export const bookApi = {
  // Get all books
  getBooks: async (): Promise<Book[]> => {
    const response = await api.get('/books');
    return response.data;
  },

  // Get book by ID
  getBookById: async (id: number): Promise<Book> => {
    const response = await api.get(`/books/${id}`);
    return response.data;
  },

  // Create new book
  createBook: async (book: CreateBookRequest): Promise<{ message: string }> => {
    const formData = new FormData();
    Object.entries(book).forEach(([key, value]) => {
      if (value !== undefined && value !== null) {
        formData.append(key, value.toString());
      }
    });

    const response = await api.post('/books', formData, {
      headers: {
        'Content-Type': 'application/x-www-form-urlencoded',
      },
    });
    return response.data;
  },

  // Update book
  updateBook: async (id: number, book: UpdateBookRequest): Promise<{ message: string }> => {
    const formData = new FormData();
    Object.entries(book).forEach(([key, value]) => {
      if (value !== undefined && value !== null) {
        formData.append(key, value.toString());
      }
    });

    const response = await api.put(`/books/${id}`, formData, {
      headers: {
        'Content-Type': 'application/x-www-form-urlencoded',
      },
    });
    return response.data;
  },

  // Delete book
  deleteBook: async (id: number): Promise<void> => {
    await api.delete(`/books/${id}`);
  },
};

export default api;