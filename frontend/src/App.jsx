// src/App.jsx
import { useEffect, useState } from "react";
import "./App.css";
// import profilePic from './agent.jpg'; // Add your profile image in src folder

function App() {
  const [agentId, setAgentId] = useState("agent001");
  const [enable, setEnable] = useState(false);

  const sendMessage = async () => {
    const phone = document.getElementById("phone").value.trim();
    const agentId = document.getElementById("agentId").value.trim();

    const res = await fetch(`http://localhost:8080/api/verify/${agentId}`);

    if (res.status === 404) {
      alert("Check your Id please.");
      return;
    } 

    if (!phone) {
      alert("Please enter a phone number.");
      return;
    }
    if (!agentId) {
      alert("Please enter agent Id.");
      return;
    }

    const formattedPhone = phone.replace(/[^0-9]/g, "");

    // Your custom message
    const message =
      encodeURIComponent(`Hello! This is a message from your agent.
      Please fill the form to proceed!

      http://localhost:8080/api/${agentId}

      `);

    // WhatsApp URL
    const whatsappUrl = `https://wa.me/${formattedPhone}?text=${message}`;

    window.open(whatsappUrl, "_blank");
  };

  return (
    <div className="body">
      <div className="inner">
        <img
          src="https://placehold.co/120x120/png"
          alt="Agent"
          className="profile"
        />
        <h1 className="title">Agent Dashboard</h1>
        {/* <h4>( {agentId} )</h4> */}
        <label htmlFor="phone">Customer Phone Number:</label>
        <input id="phone" type="text" placeholder="eg. 917050795540" />
        <label htmlFor="agentId">Agent ID:</label>
        <input id="agentId" type="text" placeholder="eg. ag1111" />

        <button onClick={sendMessage} disabled={enable}>
          Send WhatsApp Message
        </button>
      </div>
    </div>
  );
}

export default App;
