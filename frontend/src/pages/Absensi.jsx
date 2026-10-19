import { useState } from "react";

export default function Absensi() {
  const [checkedIn, setCheckedIn] = useState(false);
  const [time, setTime] = useState("");

  const handleCheckIn = () => {
    const now = new Date().toLocaleTimeString("id-ID");
    setTime(now);
    setCheckedIn(true);
  };

  const handleCheckOut = () => {
    setCheckedIn(false);
    setTime("");
  };

  return (
    <div className="max-w-md mx-auto space-y-6">
      <h1 className="text-3xl font-bold text-gray-800 text-center">Absensi</h1>
      <div className="bg-white rounded-2xl p-8 shadow-xl text-center space-y-6">
        <div className="text-6xl">{checkedIn ? "🟢" : "⚪"}</div>
        <p className="text-lg text-gray-600">
          {checkedIn ? `Check-in: ${time}` : "Belum check-in"}
        </p>
        {!checkedIn ? (
          <button onClick={handleCheckIn} className="w-full py-4 bg-green-500 hover:bg-green-600 text-white text-xl font-bold rounded-xl shadow-lg transform hover:scale-105 transition-all">
            ✅ CHECK IN
          </button>
        ) : (
          <button onClick={handleCheckOut} className="w-full py-4 bg-red-500 hover:bg-red-600 text-white text-xl font-bold rounded-xl shadow-lg transform hover:scale-105 transition-all">
            🚪 CHECK OUT
          </button>
        )}
      </div>
      <div className="bg-white rounded-xl p-6 shadow-lg">
        <h2 className="text-lg font-semibold mb-3">Riwayat Hari Ini</h2>
        <p className="text-gray-500 text-center py-4">Belum ada aktivitas</p>
      </div>
    </div>
  );
}