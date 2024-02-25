# Getting Started

## Prerequisites

Before getting started, make sure you have the following installed:
- Go programming language
- MongoDB
- API key and base URL for the currency exchange API

## Installation

### Install Go

1. Visit the [official Go website](https://golang.org/dl/) to download the Go installer for your operating system.
2. Follow the installation instructions provided on the website.

### Install MongoDB

1. Visit the [official MongoDB website](https://www.mongodb.com/try/download/community) to download the MongoDB Community Server.
2. Follow the installation instructions provided on the website.

## Setup

### Obtain API Key and Base URL

1. Sign up for an account on a currency exchange rate API provider (e.g., [Open Exchange Rates](https://openexchangerates.org/)).
2. Obtain your API key and base URL from the provider's dashboard or documentation.

### Clone the project and Install Dependencies

1. Clone the backend project repository to your local machine.
2. Navigate to the project directory.
3. Run the following command to install project dependencies:
    go mod tidy

### Configure Environment Variables

1. Create a file named `.env` in the root directory of your project.
2. Add the following lines to the `.env` file, replacing `<API_KEY>` and `<BASE_URL>` with your actual API key and base URL:

   ```plaintext
   API_KEY=<API_KEY>
   BASE_URL=<BASE_URL>
