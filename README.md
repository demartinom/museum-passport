# Museum Passport

A web application that aggregates artwork from multiple museum APIs, allowing users to search and explore art collections from institutions like the Metropolitan Museum of Art and Harvard Art Museums.

## Features

- **Multi-Museum Search**: Search artwork across multiple museum collections simultaneously
- **Detailed Artwork Pages**: View high-resolution images with comprehensive information about each piece
- **AI-Powered Summaries**: Get educational context and historical background for artworks using AI
- **Smart Caching**: Fast response times through intelligent caching of artwork data and AI summaries
- **Concurrent API Fetching**: Efficient parallel processing of multiple museum API calls

## Tech Stack

**Backend:**
- Go 1.25+
- Chi Router
- OpenAI Go SDK
- In-memory caching

**Frontend:**
- Next.js 15
- React
- TypeScript
- Tailwind CSS
- shadcn/ui components

## Getting Started

### Prerequisites

- Go 1.25 or higher
- Node.js 18+ and npm
- API keys for:
  - Harvard Art Museums ([Get key here](https://harvardartmuseums.org/collections/api))
  - OpenAI ([Get key here](https://platform.openai.com/api-keys))

### Installation

**1. Clone the repository:**
```bash
git clone https://github.com/yourusername/museum-passport.git
cd museum-passport
```

**2. Backend setup:**
```bash
cd server
cp .env.example .env
# Add your API keys to .env
go mod download
go run main.go
```

**3. Frontend setup:**
```bash
cd frontend
npm install
npm run dev
```

The backend will run on `http://localhost:3001` and frontend on `http://localhost:3000`.

### Environment Variables

**Backend (.env):**
```
HARVARD_API_KEY=your_harvard_key_here
OPENAI_API_KEY=your_openai_key_here
PORT=3001
```

**Frontend (next.config.js is already configured)**

## Usage

1. Navigate to `http://localhost:3000`
2. Use the search bar to find artwork by name, artist, or medium
3. Click on any artwork to view details and AI-generated summary
4. Use the back button to return to your search results

## API Endpoints

- `GET /api/search?general={query}&length={num}` - Search across all museums
- `GET /api/search?museum={name}&name={query}` - Search specific museum by artwork name
- `GET /api/artwork/{id}` - Get single artwork details
- `GET /api/summary?id={artworkId}` - Generate AI summary for artwork

## Project Structure
```
museum-passport/
├── server/           # Go backend
│   ├── handlers/     # HTTP handlers
│   ├── museums/      # Museum API clients
│   ├── cache/        # Caching layer
│   ├── ai/           # AI summary generation
│   └── main.go
└── frontend/         # Next.js frontend
    ├── app/          # Pages and routes
    ├── components/   # React components
    └── lib/          # API utilities
```

## Development

### Adding a New Museum API

1. Create a new client in `server/museums/`
2. Implement the `Client` interface
3. Add normalization logic for the museum's data format
4. Register the client in `main.go`

### Running Tests
```bash
cd server
go test ./...
```

## Challenges & Solutions

- **Data Normalization**: Museum APIs return inconsistent data formats. Solved by creating a unified `Artwork` struct and museum-specific normalization functions.
- **Performance**: Concurrent goroutines for parallel API calls significantly reduced response times.
- **Cost Optimization**: AI summaries are cached for 30 days to minimize OpenAI API costs.

## Future Features

- [ ] Quality scoring for better search results
- [ ] Interactive art history timeline
- [ ] Artist pages with all their works
- [ ] Similar artwork recommendations
- [ ] More museum integrations

## Blog Posts

Read about the development process:
- [Normalizing Museum API Data](your-blog-link)
- [Building AI Summaries with OpenAI](your-blog-link)

## License

MIT

## Acknowledgments

- Metropolitan Museum of Art for their open API
- Harvard Art Museums for their comprehensive API
- OpenAI for GPT-4o-mini
