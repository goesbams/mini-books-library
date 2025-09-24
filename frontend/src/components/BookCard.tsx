'use client';

import React from 'react';
import Link from 'next/link';
import { Book } from '@/types/book';
import { Calendar, User, BookOpen, Hash, Trash2, Edit, Eye } from 'lucide-react';

interface BookCardProps {
  book: Book;
  onEdit: (book: Book) => void;
  onDelete: (id: number) => void;
  onView?: (book: Book) => void;
}

export const BookCard: React.FC<BookCardProps> = ({ book, onEdit, onDelete }) => {
  const formatDate = (dateString: string) => {
    return new Date(dateString).toLocaleDateString('en-US', {
      year: 'numeric',
      month: 'long',
      day: 'numeric'
    });
  };

  return (
    <div className="bg-white rounded-lg shadow-md hover:shadow-lg transition-shadow duration-200 p-6">
      <div className="flex justify-between items-start mb-4">
        <div className="flex-1">
          <h3 className="text-xl font-semibold text-gray-900 mb-2 line-clamp-2">
            {book.title}
          </h3>
          <div className="flex items-center text-gray-600 mb-2">
            <User className="w-4 h-4 mr-2" />
            <span className="text-sm">{book.author}</span>
          </div>
        </div>
        {book.cover_image_url && (
          <div className="ml-4 flex-shrink-0">
            <img
              src={book.cover_image_url}
              alt={`${book.title} cover`}
              className="w-16 h-20 object-cover rounded border"
              onError={(e) => {
                (e.target as HTMLImageElement).style.display = 'none';
              }}
            />
          </div>
        )}
      </div>

      <div className="space-y-2 mb-4">
        <div className="flex items-center text-gray-600">
          <Calendar className="w-4 h-4 mr-2" />
          <span className="text-sm">{formatDate(book.publication_date)}</span>
        </div>
        <div className="flex items-center text-gray-600">
          <BookOpen className="w-4 h-4 mr-2" />
          <span className="text-sm">{book.number_of_pages} pages</span>
        </div>
        <div className="flex items-center text-gray-600">
          <Hash className="w-4 h-4 mr-2" />
          <span className="text-sm font-mono">{book.isbn}</span>
        </div>
      </div>

      {book.description && (
        <p className="text-gray-700 text-sm mb-4 line-clamp-3">
          {book.description}
        </p>
      )}

      <div className="flex justify-end space-x-2">
        <Link
          href={`/books/${book.id}`}
          className="flex items-center px-3 py-1.5 text-sm text-blue-600 hover:text-blue-800 hover:bg-blue-50 rounded transition-colors"
        >
          <Eye className="w-4 h-4 mr-1" />
          View
        </Link>
        <button
          onClick={() => onEdit(book)}
          className="flex items-center px-3 py-1.5 text-sm text-green-600 hover:text-green-800 hover:bg-green-50 rounded transition-colors"
        >
          <Edit className="w-4 h-4 mr-1" />
          Edit
        </button>
        <button
          onClick={() => onDelete(book.id)}
          className="flex items-center px-3 py-1.5 text-sm text-red-600 hover:text-red-800 hover:bg-red-50 rounded transition-colors"
        >
          <Trash2 className="w-4 h-4 mr-1" />
          Delete
        </button>
      </div>
    </div>
  );
};