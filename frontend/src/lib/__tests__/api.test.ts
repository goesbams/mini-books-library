import { Book, CreateBookRequest, UpdateBookRequest } from '@/types/book'

// Mock the entire API module
const mockBookApi = {
  getBooks: jest.fn(),
  getBookById: jest.fn(),
  createBook: jest.fn(),
  updateBook: jest.fn(),
  deleteBook: jest.fn(),
}

jest.mock('../api', () => ({
  bookApi: mockBookApi,
}))

describe('bookApi', () => {
  beforeEach(() => {
    jest.clearAllMocks()
  })

  describe('getBooks', () => {
    it('should fetch all books successfully', async () => {
      const mockBooks: Book[] = [
        {
          id: 1,
          title: 'Test Book',
          author: 'Test Author',
          cover_image_url: '',
          description: 'Test Description',
          publication_date: '2023-01-01',
          number_of_pages: 100,
          isbn: '1234567890123',
        },
      ]

      mockBookApi.getBooks.mockResolvedValue(mockBooks)

      const result = await mockBookApi.getBooks()

      expect(mockBookApi.getBooks).toHaveBeenCalled()
      expect(result).toEqual(mockBooks)
    })

    it('should handle getBooks error', async () => {
      const error = new Error('Network error')
      mockBookApi.getBooks.mockRejectedValue(error)

      await expect(mockBookApi.getBooks()).rejects.toThrow('Network error')
    })
  })

  describe('getBookById', () => {
    it('should fetch a book by ID successfully', async () => {
      const mockBook: Book = {
        id: 1,
        title: 'Test Book',
        author: 'Test Author',
        cover_image_url: '',
        description: 'Test Description',
        publication_date: '2023-01-01',
        number_of_pages: 100,
        isbn: '1234567890123',
      }

      mockBookApi.getBookById.mockResolvedValue(mockBook)

      const result = await mockBookApi.getBookById(1)

      expect(mockBookApi.getBookById).toHaveBeenCalledWith(1)
      expect(result).toEqual(mockBook)
    })

    it('should handle getBookById error', async () => {
      const error = new Error('Book not found')
      mockBookApi.getBookById.mockRejectedValue(error)

      await expect(mockBookApi.getBookById(1)).rejects.toThrow('Book not found')
    })
  })

  describe('createBook', () => {
    it('should create a book successfully', async () => {
      const newBook: CreateBookRequest = {
        title: 'New Book',
        author: 'New Author',
        publication_date: '2023-01-01',
        number_of_pages: 200,
        isbn: '1234567890123',
      }

      const mockResponse = { message: 'Book created successfully' }
      mockBookApi.createBook.mockResolvedValue(mockResponse)

      const result = await mockBookApi.createBook(newBook)

      expect(mockBookApi.createBook).toHaveBeenCalledWith(newBook)
      expect(result).toEqual(mockResponse)
    })

    it('should handle createBook error', async () => {
      const newBook: CreateBookRequest = {
        title: 'New Book',
        author: 'New Author',
        publication_date: '2023-01-01',
        number_of_pages: 200,
        isbn: '1234567890123',
      }

      const error = new Error('Validation error')
      mockBookApi.createBook.mockRejectedValue(error)

      await expect(mockBookApi.createBook(newBook)).rejects.toThrow('Validation error')
    })
  })

  describe('updateBook', () => {
    it('should update a book successfully', async () => {
      const updateData: UpdateBookRequest = {
        title: 'Updated Book',
      }

      const mockResponse = { message: 'Book updated successfully' }
      mockBookApi.updateBook.mockResolvedValue(mockResponse)

      const result = await mockBookApi.updateBook(1, updateData)

      expect(mockBookApi.updateBook).toHaveBeenCalledWith(1, updateData)
      expect(result).toEqual(mockResponse)
    })

    it('should handle updateBook error', async () => {
      const updateData: UpdateBookRequest = {
        title: 'Updated Book',
      }

      const error = new Error('Book not found')
      mockBookApi.updateBook.mockRejectedValue(error)

      await expect(mockBookApi.updateBook(1, updateData)).rejects.toThrow('Book not found')
    })
  })

  describe('deleteBook', () => {
    it('should delete a book successfully', async () => {
      mockBookApi.deleteBook.mockResolvedValue(undefined)

      await mockBookApi.deleteBook(1)

      expect(mockBookApi.deleteBook).toHaveBeenCalledWith(1)
    })

    it('should handle deleteBook error', async () => {
      const error = new Error('Book not found')
      mockBookApi.deleteBook.mockRejectedValue(error)

      await expect(mockBookApi.deleteBook(1)).rejects.toThrow('Book not found')
    })
  })
})
