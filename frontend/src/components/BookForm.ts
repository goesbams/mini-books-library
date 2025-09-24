'use client';

import React from 'react';
import { useForm } from 'react-hook-form';
import { zodResolver } from '@hookform/resolvers/zod';
import { bookSchema, BookFormData } from '@/lib/validations';
import { useBooks } from '@/contexts/BookContext';
import { Book } from '@/types/book';
import { X, Save, Loader2 } from 'lucide-react';

interface BookFormProps {
  isOpen: boolean;
  onClose: () => void;
  onSuccess: () => void;
  book?: Book;
}

export const BookForm: React.FC<BookFormProps> = ({ isOpen, onClose, onSuccess, book }) => {
  const { createBook, updateBook, state } = useBooks();
  const isEditing = !!book;

  const {
    register,
    handleSubmit,
    formState: { errors, isSubmitting },
    reset,
    setValue
  } = useForm<BookFormData>({
    resolver: zodResolver(bookSchema),
    defaultValues: {
      title: book?.title || '',
      author: book?.author || '',
      cover_image_url: book?.cover_image_url || '',
      description: book?.description || '',
      publication_date: book?.publication_date || '',
      number_of_pages: book?.number_of_pages || 0,
      isbn: book?.isbn || '',
    }
  });

  React.useEffect(() => {
    if (book) {
      setValue('title', book.title);
      setValue('author', book.author);
      setValue('cover_image_url', book.cover_image_url);
      setValue('description', book.description);
      setValue('publication_date', book.publication_date);
      setValue('number_of_pages', book.number_of_pages);
      setValue('isbn', book.isbn);
    }
  }, [book, setValue]);

  const onSubmit = async (data: BookFormData) => {
    try {
      if (isEditing && book) {
        await updateBook(book.id, data);
      } else {
        await createBook(data);
      }
      reset();
      onSuccess();
    } catch {
      // Error is handled by the context
    }
  };

  const handleClose = () => {
    reset();
    onClose();
  };

  if (!isOpen) return null;

  return (
    <div className="fixed inset-0 bg-gray-600 bg-opacity-50 overflow-y-auto h-full w-full z-50">
      <div className="relative top-20 mx-auto p-5 border w-full max-w-2xl shadow-lg rounded-md bg-white">
        <div className="flex justify-between items-center mb-6">
          <h3 className="text-lg font-medium text-gray-900">
            {isEditing ? 'Edit Book' : 'Add New Book'}
          </h3>
          <button
            onClick={handleClose}
            className="text-gray-400 hover:text-gray-600"
          >
            <X className="w-6 h-6" />
          </button>
        </div>

        <form onSubmit={handleSubmit(onSubmit)} className="space-y-6">
          <div className="grid grid-cols-1 md:grid-cols-2 gap-6">
            {/* Title */}
            <div className="md:col-span-2">
              <label htmlFor="title" className="block text-sm font-medium text-gray-700">
                Title *
              </label>
              <input
                {...register('title')}
                type="text"
                className="mt-1 block w-full border-gray-300 rounded-md shadow-sm focus:ring-blue-500 focus:border-blue-500 sm:text-sm"
                placeholder="Enter book title"
              />
              {errors.title && (
                <p className="mt-1 text-sm text-red-600">{errors.title.message}</p>
              )}
            </div>

            {/* Author */}
            <div className="md:col-span-2">
              <label htmlFor="author" className="block text-sm font-medium text-gray-700">
                Author *
              </label>
              <input
                {...register('author')}
                type="text"
                className="mt-1 block w-full border-gray-300 rounded-md shadow-sm focus:ring-blue-500 focus:border-blue-500 sm:text-sm"
                placeholder="Enter author name"
              />
              {errors.author && (
                <p className="mt-1 text-sm text-red-600">{errors.author.message}</p>
              )}
            </div>

            {/* Publication Date */}
            <div>
              <label htmlFor="publication_date" className="block text-sm font-medium text-gray-700">
                Publication Date *
              </label>
              <input
                {...register('publication_date')}
                type="date"
                className="mt-1 block w-full border-gray-300 rounded-md shadow-sm focus:ring-blue-500 focus:border-blue-500 sm:text-sm"
              />
              {errors.publication_date && (
                <p className="mt-1 text-sm text-red-600">{errors.publication_date.message}</p>
              )}
            </div>

            {/* Number of Pages */}
            <div>
              <label htmlFor="number_of_pages" className="block text-sm font-medium text-gray-700">
                Number of Pages *
              </label>
              <input
                {...register('number_of_pages', { valueAsNumber: true })}
                type="number"
                min="1"
                className="mt-1 block w-full border-gray-300 rounded-md shadow-sm focus:ring-blue-500 focus:border-blue-500 sm:text-sm"
                placeholder="Enter number of pages"
              />
              {errors.number_of_pages && (
                <p className="mt-1 text-sm text-red-600">{errors.number_of_pages.message}</p>
              )}
            </div>

            {/* ISBN */}
            <div>
              <label htmlFor="isbn" className="block text-sm font-medium text-gray-700">
                ISBN *
              </label>
              <input
                {...register('isbn')}
                type="text"
                maxLength={13}
                className="mt-1 block w-full border-gray-300 rounded-md shadow-sm focus:ring-blue-500 focus:border-blue-500 sm:text-sm"
                placeholder="Enter 13-digit ISBN"
              />
              {errors.isbn && (
                <p className="mt-1 text-sm text-red-600">{errors.isbn.message}</p>
              )}
            </div>

            {/* Cover Image URL */}
            <div>
              <label htmlFor="cover_image_url" className="block text-sm font-medium text-gray-700">
                Cover Image URL
              </label>
              <input
                {...register('cover_image_url')}
                type="url"
                className="mt-1 block w-full border-gray-300 rounded-md shadow-sm focus:ring-blue-500 focus:border-blue-500 sm:text-sm"
                placeholder="Enter cover image URL"
              />
              {errors.cover_image_url && (
                <p className="mt-1 text-sm text-red-600">{errors.cover_image_url.message}</p>
              )}
            </div>
          </div>

          {/* Description */}
          <div>
            <label htmlFor="description" className="block text-sm font-medium text-gray-700">
              Description
            </label>
            <textarea
              {...register('description')}
              rows={4}
              className="mt-1 block w-full border-gray-300 rounded-md shadow-sm focus:ring-blue-500 focus:border-blue-500 sm:text-sm"
              placeholder="Enter book description"
            />
            {errors.description && (
              <p className="mt-1 text-sm text-red-600">{errors.description.message}</p>
            )}
          </div>

          {/* Form Actions */}
          <div className="flex justify-end space-x-3 pt-6 border-t">
            <button
              type="button"
              onClick={handleClose}
              className="px-4 py-2 text-sm font-medium text-gray-700 bg-white border border-gray-300 rounded-md hover:bg-gray-50 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500"
            >
              Cancel
            </button>
            <button
              type="submit"
              disabled={isSubmitting || state.loading}
              className="flex items-center px-4 py-2 text-sm font-medium text-white bg-blue-600 border border-transparent rounded-md hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500 disabled:opacity-50"
            >
              {isSubmitting || state.loading ? (
                <Loader2 className="w-4 h-4 mr-2 animate-spin" />
              ) : (
                <Save className="w-4 h-4 mr-2" />
              )}
              {isEditing ? 'Update Book' : 'Add Book'}
            </button>
          </div>
        </form>
      </div>
    </div>
  );
};