'use client';

import React, { createContext, useContext, useReducer, useEffect, ReactNode } from 'react';
import { Book, CreateBookRequest, UpdateBookRequest } from '@/types/book';
import { bookApi } from '@/lib/api';
import { AxiosError } from 'axios';

interface BookState {
  books: Book[];
  loading: boolean;
  error: string | null;
}

type BookAction =
  | { type: 'SET_LOADING'; payload: boolean }
  | { type: 'SET_BOOKS'; payload: Book[] }
  | { type: 'ADD_BOOK'; payload: Book }
  | { type: 'UPDATE_BOOK'; payload: Book }
  | { type: 'DELETE_BOOK'; payload: number }
  | { type: 'SET_ERROR'; payload: string | null };

interface BookContextType {
  state: BookState;
  fetchBooks: () => Promise<void>;
  createBook: (book: CreateBookRequest) => Promise<void>;
  updateBook: (id: number, book: UpdateBookRequest) => Promise<void>;
  deleteBook: (id: number) => Promise<void>;
  clearError: () => void;
}

const BookContext = createContext<BookContextType | undefined>(undefined);

// Helper function to extract error message
const getErrorMessage = (error: unknown, defaultMessage: string): string => {
  if (error instanceof AxiosError) {
    return error.response?.data?.message || defaultMessage;
  }
  if (error instanceof Error) {
    return error.message;
  }
  return defaultMessage;
};

const bookReducer = (state: BookState, action: BookAction): BookState => {
  switch (action.type) {
    case 'SET_LOADING':
      return { ...state, loading: action.payload };
    case 'SET_BOOKS':
      return { ...state, books: action.payload, loading: false, error: null };
    case 'ADD_BOOK':
      return { ...state, books: [...state.books, action.payload], loading: false, error: null };
    case 'UPDATE_BOOK':
      return {
        ...state,
        books: state.books.map(book => 
          book.id === action.payload.id ? action.payload : book
        ),
        loading: false,
        error: null
      };
    case 'DELETE_BOOK':
      return {
        ...state,
        books: state.books.filter(book => book.id !== action.payload),
        loading: false,
        error: null
      };
    case 'SET_ERROR':
      return { ...state, error: action.payload, loading: false };
    default:
      return state;
  }
};

const initialState: BookState = {
  books: [],
  loading: false,
  error: null,
};

export const BookProvider: React.FC<{ children: ReactNode }> = ({ children }) => {
  const [state, dispatch] = useReducer(bookReducer, initialState);

  const fetchBooks = async () => {
    try {
      dispatch({ type: 'SET_LOADING', payload: true });
      const books = await bookApi.getBooks();
      dispatch({ type: 'SET_BOOKS', payload: books });
    } catch (error: unknown) {
      const errorMessage = getErrorMessage(error, 'Failed to fetch books');
      dispatch({ 
        type: 'SET_ERROR', 
        payload: errorMessage
      });
    }
  };

  const createBook = async (book: CreateBookRequest) => {
    try {
      dispatch({ type: 'SET_LOADING', payload: true });
      await bookApi.createBook(book);
      // Refresh the books list after creating
      await fetchBooks();
    } catch (error: unknown) {
      const errorMessage = getErrorMessage(error, 'Failed to create book');
      dispatch({ 
        type: 'SET_ERROR', 
        payload: errorMessage
      });
    }
  };

  const updateBook = async (id: number, book: UpdateBookRequest) => {
    try {
      dispatch({ type: 'SET_LOADING', payload: true });
      await bookApi.updateBook(id, book);
      // Refresh the books list after updating
      await fetchBooks();
    } catch (error: unknown) {
      const errorMessage = getErrorMessage(error, 'Failed to update book');
      dispatch({ 
        type: 'SET_ERROR', 
        payload: errorMessage
      });
    }
  };

  const deleteBook = async (id: number) => {
    try {
      dispatch({ type: 'SET_LOADING', payload: true });
      await bookApi.deleteBook(id);
      dispatch({ type: 'DELETE_BOOK', payload: id });
    } catch (error: unknown) {
      const errorMessage = getErrorMessage(error, 'Failed to delete book');
      dispatch({ 
        type: 'SET_ERROR', 
        payload: errorMessage
      });
    }
  };

  const clearError = () => {
    dispatch({ type: 'SET_ERROR', payload: null });
  };

  // Fetch books on component mount
  useEffect(() => {
    fetchBooks();
  }, []);

  const value: BookContextType = {
    state,
    fetchBooks,
    createBook,
    updateBook,
    deleteBook,
    clearError,
  };

  return (
    <BookContext.Provider value={value}>
      {children}
    </BookContext.Provider>
  );
};

export const useBooks = (): BookContextType => {
  const context = useContext(BookContext);
  if (context === undefined) {
    throw new Error('useBooks must be used within a BookProvider');
  }
  return context;
};