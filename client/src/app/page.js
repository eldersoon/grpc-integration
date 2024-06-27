"use client";
import { useState } from "react";
import { TrezoitaoServiceClient } from "../38tao_grpc_web_pb";
import { GiveMeBTCRequest } from "../38tao_pb";

const client = new TrezoitaoServiceClient("http://localhost:8080");

export default function Home() {
  const [key, setKey] = useState("");
  const [amount, setAmount] = useState(0);
  const [reply, setReply] = useState("");

  const sendRequest = async () => {
    const request = new GiveMeBTCRequest();
    request.setKey(key);
    request.setAmount(amount);

    await client.getBTC(request, {}, (err, response) => {
      console.log("????");
      if (err) {
        console.error(err);
      } else {
        setReply(response.getMessage());
      }
    });
    console.log("????");
  };

  return (
    <div>
      <div className="flex flex-row gap-[16px]">
        Chave PIX:
        <input
          type="text"
          value={key}
          onChange={(e) => setKey(e.target.value)}
          placeholder="Key"
        />
      </div>
      <div className="flex flex-row gap-[16px]">
        Valor:
        <input
          type="number"
          value={amount}
          onChange={(e) => setAmount(parseInt(e.target.value, 10))}
          placeholder="Amount"
        />
      </div>
      <button
        className="bg-[#cecece] px-[16px] py-[8px] rounded-sm"
        onClick={sendRequest}
      >
        Enviar
      </button>
      <p>Mensagem: {reply}</p>
    </div>
  );
}
