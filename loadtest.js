import http from 'k6/http';
import { check, sleep } from 'k6';

// Load testing configurations
export const options = {
  stages: [
    { duration: '30s', target: 10 }, // Ramp-up to 10 VUs over 30 seconds
    { duration: '10m', target: 100 }, // Stay at 50 VUs for 1 minute
    { duration: '30s', target: 0 }, // Ramp-down to 0 VUs over 30 seconds
  ],
};

// Base URL of your API
let BASE_URL;
if (__ENV.ENVIRONMENT == "production") {
  BASE_URL = "https://api.clarkkiee.dev"
} else {
  BASE_URL = "http://localhost:8080"
}

// Example token (replace with a valid token from your API)
const TOKEN = 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6IjhhN2M1N2JmLTJkYTMtNGJhNC1iOGI1LWY1NWFkYTRmZTM5OSIsImlzcyI6IlRlbXBsYXRlIiwiZXhwIjoxNzQ2MjEyNDM5LCJpYXQiOjE3NDYyMDg4Mzl9.lRywBlLIhnhFsZoB5VnXK_KXCwYe44HgJWiK-kmQFYU'; // Replace with a valid JWT from your application

// List of endpoints to test
const endpoints = [
  { method: 'GET', path: '/api/store/me', secured: true },
  { method: 'POST', path: '/api/store/register', body: { name: "Store ABC", email: "store@test.com", password: "password123" } },
  { method: 'POST', path: '/api/store/login', body: { email: "store@test.com", password: "password123" } },
  { method: 'POST', path: '/api/address/new', body: { street: "123 Main St", city: "Metropolis", postal_code: "12345" }, secured: true },
  { method: 'GET', path: '/api/address/all', secured: true },
  { method: 'POST', path: '/api/product/create', body: { name: "Sample Product", price: 100.0, stock: 10 }, secured: true },
  { method: 'GET', path: '/api/product/store', secured: true },
  { method: 'GET', path: '/api/product/public?limit=10&offset=0&distance=10000' },
  { method: 'GET', path: '/api/order/', secured: true },
  { method: 'POST', path: '/api/order/add', body: { product_id: 1, quantity: 2 }, secured: true },
];

// Main load testing function
export default function () {
  // Randomly pick an endpoint to test
  const endpoint = endpoints[Math.floor(Math.random() * endpoints.length)];
  const url = `${BASE_URL}${endpoint.path}`;

  // Set request headers
  const headers = { 'Content-Type': 'application/json' };
  if (endpoint.secured) {
    headers['Authorization'] = `Bearer ${TOKEN}`; 
  }

  // Execute request based on method
  let res;
  if (endpoint.method === 'GET') {
    res = http.get(url, { headers });
  } else if (endpoint.method === 'POST') {
    res = http.post(url, JSON.stringify(endpoint.body), { headers });
  }

  // Validate response
  check(res, {
    'status is 200': (r) => r.status === 200,
    'response time is below 500ms': (r) => r.timings.duration < 500,
  });

  // Pause for 1 second between requests
  sleep(1);
}