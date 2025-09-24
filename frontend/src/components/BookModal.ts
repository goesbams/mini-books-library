'use client';

import React from 'react';
import { Book } from '@/types/book';
import { X, Calendar, User, BookOpen, Hash, ExternalLink } from 'lucide-react';

interface BookModalProps {
  isOpen: boolean;
  onClose: () => void;
  book: Book;
}

export const BookModal: React.FC<BookModalProps> = ({ isOpen, onClose, book }) => {
  if (!isOpen) return null;

  const formatDate = (dateString: string) => {
    return new Date(dateString).toLocaleDateString('en-US', {
      year: 'numeric',
      month: 'long',
      day: 'numeric'
    });
  };

  return (
    <div className="fixed inset-0 bg-gray-600 bg-opacity-50 overflow-y-auto h-full w-full z-50">
      <div className="relative top-20 mx-auto p-5 border w-full max-w-2xl shadow-lg rounded-md bg-white">
        <div className="flex justify-between items-start mb-6">
          <div className="flex-1">
            <h2 className="text-2xl font-bold text-gray-900 mb-2">{book.title}</h2>
            <div className="flex items-center text-gray-600 mb-4">
              <User className="w-5 h-5 mr-2" />
              <span className="text-lg">{book.author}</span>
            </div>
          </div>
          <button
            onClick={onClose}
            className="text-gray-400 hover:text-gray-600 ml-4"
          >
            <X className="w-6 h-6" />
          </button>
        </div>

        <div className="grid grid-cols-1 md:grid-cols-3 gap-6">
          {/* Book Cover */}
          <div className="md:col-span-1">
            {book.cover_image_url ? (
              <img
                src={book.cover_image_url}
                alt={`${book.title} cover`}
                className="w-full h-64 object-cover rounded-lg border shadow-md"
                onError={(e) => {
                  (e.target as HTMLImageElement).src = '/placeholder-book.png';
                }}
              />
            ) : (
              <div className="w-full h-64 bg-gray-200 rounded-lg border shadow-md flex items-center justify-center">
                <BookOpen className="w-16 h-16 text-gray-400" />
              </div>
            )}
          </div>

          {/* Book Details */}
          <div className="md:col-span-2 space-y-4">
            <div className="grid grid-cols-1 sm:grid-cols-2 gap-4">
              <div className="flex items-center text-gray-600">
                <Calendar className="w-5 h-5 mr-3" />
                <div>
                  <p className="text-sm font-medium text-gray-500">Publication Date</p>
                  <p className="text-sm">{formatDate(book.publication_date)}</p>
                </div>
              </div>

              <div className="flex items-center text-gray-600">
                <BookOpen className="w-5 h-5 mr-3" />
                <div>
                  <p className="text-sm font-medium text-gray-500">Pages</p>
                  <p className="text-sm">{book.number_of_pages}</p>
                </div>
              </div>

              <div className="flex items-center text-gray-600 sm:col-span-2">
                <Hash className="w-5 h-5 mr-3" />
                <div>
                  <p className="text-sm font-medium text-gray-500">ISBN</p>
                  <p className="text-sm font-mono">{book.isbn}</p>
                </div>
              </div>
            </div>

            {book.description && (
              <div>
                <h3 className="text-lg font-semibold text-gray-900 mb-2">Description</h3>
                <p className="text-gray-700 leading-relaxed">{book.description}</p>
              </div>
            )}

            {book.cover_image_url && (
              <div>
                <a
                  href={book.cover_image_url}
                  target="_blank"
                  rel="noopener noreferrer"
                  className="inline-flex items-center text-blue-600 hover:text-blue-800 text-sm"
                >
                  <ExternalLink className="w-4 h-4 mr-1" />
                  View full cover image
                </a>
              </div>
            )}
          </div>
        </div>

        {/* Modal Actions */}
        <div className="flex justify-end pt-6 border-t mt-6">
          <button
            onClick={onClose}
            className="px-4 py-2 text-sm font-medium text-gray-700 bg-white border border-gray-300 rounded-md hover:bg-gray-50 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500"
          >
            Close
          </button>
        </div>
      </div>
    </div>
  );
};