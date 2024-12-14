const API_BASE_URL = "http://localhost:8080"; // Replace with your backend's URL

export async function fetchPolls() {
  const response = await fetch(`${API_BASE_URL}/polls`);
  return response.json();
}

export async function createPoll(poll) {
  const response = await fetch(`${API_BASE_URL}/polls`, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify(poll),
  });
  return response.json();
}

export async function vote(optionId) {
  const response = await fetch(`${API_BASE_URL}/options/${optionId}/vote`, {
    method: "POST",
  });
  return response.json();
}
