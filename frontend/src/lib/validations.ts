import { z } from 'zod';

export const bookSchema = z.object({
  title: z.string()
    .min(2, 'Title must be at least 2 characters')
    .max(255, 'Title must be less than 255 characters'),
  author: z.string()
    .min(2, 'Author must be at least 2 characters')
    .max(255, 'Author must be less than 255 characters'),
  cover_image_url: z.string()
    .url('Must be a valid URL')
    .optional()
    .or(z.literal('')),
  description: z.string()
    .max(1000, 'Description must be less than 1000 characters')
    .optional()
    .or(z.literal('')),
  publication_date: z.string()
    .regex(/^\d{4}-\d{2}-\d{2}$/, 'Date must be in YYYY-MM-DD format'),
  number_of_pages: z.number()
    .min(1, 'Number of pages must be at least 1')
    .int('Number of pages must be an integer'),
  isbn: z.string()
    .length(13, 'ISBN must be exactly 13 digits')
    .regex(/^\d{13}$/, 'ISBN must contain only numbers'),
});

export const updateBookSchema = bookSchema.partial();

export type BookFormData = z.infer<typeof bookSchema>;
export type UpdateBookFormData = z.infer<typeof updateBookSchema>;