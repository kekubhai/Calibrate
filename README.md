


Calibrate
 - Real-time WebSocket Messaging System
A Go-based application that demonstrates real-time communication using WebSockets with separate client and server components. The application also integrates with Finnhub API to access financial market data.

Features
Real-time bidirectional communication between client and server using WebSockets
Terminal-based client interface for sending and receiving messages
Server timestamp appended to each response
Integration with Finnhub API for financial data (stocks and crypto)
Support for subscribing to multiple financial symbols (AAPL, AMZN, BINANCE:BTCUSDT)
Project Structure
Prerequisites
Go 1.x or later
Finnhub API token (configured in .env file)
Installation
Clone the repository:
Install dependencies:
Running the Application
Start the server:
In a separate terminal, run the client:
Usage
After starting both server and client, the client will connect to the server automatically
Type messages in the client terminal and press Enter to send
The server will receive your message, append a timestamp, and send it back
The client will display the response from the server
Finnhub Integration
The application can connect to Finnhub's WebSocket API to receive real-time financial data. To use this feature:

Make sure your Finnhub API token is set in the .env file
Call the connectToFinnhub() function from your code
Technologies
Go - Programming language
Gorilla WebSocket - WebSocket implementation for Go
Finnhub API - Financial market data API