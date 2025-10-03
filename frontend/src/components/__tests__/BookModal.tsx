import React from 'react'
import { render, screen } from '@testing-library/react'
import userEvent from '@testing-library/user-event'
import { BookModal } from '../BookModal'
import { Book } from '@/types/book'

const mockBook: Book = {
  id: 1,
  title: 'Test Book Title',
  author: 'Test Author',
  cover_image_url: 'https://example.com/cover.jpg',
  description: 'This is a comprehensive test book description that provides detailed information about the book content and its significance.',
  publication_date: '2023-01-15',
  number_of_pages: 250,
  isbn: '1234567890123',
}

const mockHandlers = {
  onClose: jest.fn(),
}

describe('BookModal', () => {
  beforeEach(() => {
    jest.clearAllMocks()
  })

  it('should not render when isOpen is false', () => {
    render(<BookModal isOpen={false} book={mockBook} {...mockHandlers} />)
    
    expect(screen.queryByText('Test Book Title')).not.toBeInTheDocument()
  })

  it('should render book information when isOpen is true', () => {
    render(<BookModal isOpen={true} book={mockBook} {...mockHandlers} />)
    
    expect(screen.getByText('Test Book Title')).toBeInTheDocument()
    expect(screen.getByText('Test Author')).toBeInTheDocument()
    expect(screen.getByText('January 15, 2023')).toBeInTheDocument()
    expect(screen.getByText('250')).toBeInTheDocument()
    expect(screen.getByText('1234567890123')).toBeInTheDocument()
    expect(screen.getByText('This is a comprehensive test book description that provides detailed information about the book content and its significance.')).toBeInTheDocument()
  })

  it('should render cover image when provided', () => {
    render(<BookModal isOpen={true} book={mockBook} {...mockHandlers} />)
    
    const coverImage = screen.getByAltText('Test Book Title cover')
    expect(coverImage).toBeInTheDocument()
    expect(coverImage).toHaveAttribute('src', 'https://example.com/cover.jpg')
  })

  it('should render placeholder when no cover image is provided', () => {
    const bookWithoutCover = { ...mockBook, cover_image_url: '' }
    render(<BookModal isOpen={true} book={bookWithoutCover} {...mockHandlers} />)
    
    // Should show placeholder icon instead of image
    expect(screen.queryByAltText('Test Book Title cover')).not.toBeInTheDocument()
  })

  it('should call onClose when close button is clicked', async () => {
    const user = userEvent.setup()
    render(<BookModal isOpen={true} book={mockBook} {...mockHandlers} />)
    
    // Find the close button by its SVG icon (X)
    const closeButton = screen.getByRole('button', { name: '' })
    await user.click(closeButton)
    
    expect(mockHandlers.onClose).toHaveBeenCalled()
  })

  it('should call onClose when modal close button is clicked', async () => {
    const user = userEvent.setup()
    render(<BookModal isOpen={true} book={mockBook} {...mockHandlers} />)
    
    const modalCloseButton = screen.getByText('Close')
    await user.click(modalCloseButton)
    
    expect(mockHandlers.onClose).toHaveBeenCalled()
  })

  it('should format date correctly', () => {
    const bookWithDifferentDate = {
      ...mockBook,
      publication_date: '2022-12-25',
    }
    render(<BookModal isOpen={true} book={bookWithDifferentDate} {...mockHandlers} />)
    
    expect(screen.getByText('December 25, 2022')).toBeInTheDocument()
  })

  it('should display ISBN in monospace font', () => {
    render(<BookModal isOpen={true} book={mockBook} {...mockHandlers} />)
    
    const isbnElement = screen.getByText('1234567890123')
    expect(isbnElement).toHaveClass('font-mono')
  })

  it('should show external link for cover image when URL is provided', () => {
    render(<BookModal isOpen={true} book={mockBook} {...mockHandlers} />)
    
    const externalLink = screen.getByText('View full cover image')
    expect(externalLink).toBeInTheDocument()
    expect(externalLink.closest('a')).toHaveAttribute('href', 'https://example.com/cover.jpg')
    expect(externalLink.closest('a')).toHaveAttribute('target', '_blank')
    expect(externalLink.closest('a')).toHaveAttribute('rel', 'noopener noreferrer')
  })

  it('should not show external link when no cover image URL is provided', () => {
    const bookWithoutCover = { ...mockBook, cover_image_url: '' }
    render(<BookModal isOpen={true} book={bookWithoutCover} {...mockHandlers} />)
    
    expect(screen.queryByText('View full cover image')).not.toBeInTheDocument()
  })

  it('should handle book without description', () => {
    const bookWithoutDescription = { ...mockBook, description: '' }
    render(<BookModal isOpen={true} book={bookWithoutDescription} {...mockHandlers} />)
    
    expect(screen.getByText('Test Book Title')).toBeInTheDocument()
    expect(screen.getByText('Test Author')).toBeInTheDocument()
    // Description section should not be rendered
    expect(screen.queryByText('Description')).not.toBeInTheDocument()
  })

  it('should display all book metadata correctly', () => {
    render(<BookModal isOpen={true} book={mockBook} {...mockHandlers} />)
    
    // Check all metadata labels
    expect(screen.getByText('Publication Date')).toBeInTheDocument()
    expect(screen.getByText('Pages')).toBeInTheDocument()
    expect(screen.getByText('ISBN')).toBeInTheDocument()
    expect(screen.getByText('Description')).toBeInTheDocument()
  })
})
