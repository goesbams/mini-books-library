import React from 'react'
import { render, screen } from '@testing-library/react'
import userEvent from '@testing-library/user-event'
import { BookCard } from '../BookCard'
import { Book } from '@/types/book'

const mockBook: Book = {
  id: 1,
  title: 'Test Book Title',
  author: 'Test Author',
  cover_image_url: 'https://example.com/cover.jpg',
  description: 'This is a test book description that provides details about the book content.',
  publication_date: '2023-01-15',
  number_of_pages: 250,
  isbn: '1234567890123',
}

const mockHandlers = {
  onEdit: jest.fn(),
  onDelete: jest.fn(),
}

describe('BookCard', () => {
  beforeEach(() => {
    jest.clearAllMocks()
  })

  it('should render book information correctly', () => {
    render(<BookCard book={mockBook} {...mockHandlers} />)

    expect(screen.getByText('Test Book Title')).toBeInTheDocument()
    expect(screen.getByText('Test Author')).toBeInTheDocument()
    expect(screen.getByText('January 15, 2023')).toBeInTheDocument()
    expect(screen.getByText('250 pages')).toBeInTheDocument()
    expect(screen.getByText('1234567890123')).toBeInTheDocument()
    expect(screen.getByText('This is a test book description that provides details about the book content.')).toBeInTheDocument()
  })

  it('should render cover image when provided', () => {
    render(<BookCard book={mockBook} {...mockHandlers} />)

    const coverImage = screen.getByAltText('Test Book Title cover')
    expect(coverImage).toBeInTheDocument()
    expect(coverImage).toHaveAttribute('src', 'https://example.com/cover.jpg')
  })

  it('should not render cover image when not provided', () => {
    const bookWithoutCover = { ...mockBook, cover_image_url: '' }
    render(<BookCard book={bookWithoutCover} {...mockHandlers} />)

    expect(screen.queryByAltText('Test Book Title cover')).not.toBeInTheDocument()
  })

  it('should call onEdit when edit button is clicked', async () => {
    const user = userEvent.setup()
    render(<BookCard book={mockBook} {...mockHandlers} />)

    const editButton = screen.getByText('Edit')
    await user.click(editButton)

    expect(mockHandlers.onEdit).toHaveBeenCalledWith(mockBook)
  })

  it('should call onDelete when delete button is clicked', async () => {
    const user = userEvent.setup()
    render(<BookCard book={mockBook} {...mockHandlers} />)

    const deleteButton = screen.getByText('Delete')
    await user.click(deleteButton)

    expect(mockHandlers.onDelete).toHaveBeenCalledWith(1)
  })

  it('should have a working view link', () => {
    render(<BookCard book={mockBook} {...mockHandlers} />)

    const viewLink = screen.getByText('View').closest('a')
    expect(viewLink).toHaveAttribute('href', '/books/1')
  })

  it('should format date correctly', () => {
    const bookWithDifferentDate = {
      ...mockBook,
      publication_date: '2022-12-25',
    }
    render(<BookCard book={bookWithDifferentDate} {...mockHandlers} />)

    expect(screen.getByText('December 25, 2022')).toBeInTheDocument()
  })

  it('should display ISBN in monospace font', () => {
    render(<BookCard book={mockBook} {...mockHandlers} />)

    const isbnElement = screen.getByText('1234567890123')
    expect(isbnElement).toHaveClass('font-mono')
  })
})
