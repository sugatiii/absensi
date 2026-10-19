import { useState, useEffect } from "react";

export default function Dashboard() {
  const [stats] = useState({
    totalKaryawan: 45,
    hadirHariIni: 38,
    terlambat: 3,
    izin: 4,
  });

  return (
    <div className="space-y-6">
      <h1 className="text-3xl font-bold text-gray-800">Dashboard Admin</h1>
      <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-4">
        {[
          { label: "Total Karyawan", value: stats.totalKaryawan, color: "bg-blue-500" },
          { label: "Hadir Hari Ini", value: stats.hadirHariIni, color: "bg-green-500" },
          { label: "Terlambat", value: stats.terlambat, color: "bg-yellow-500" },
          { label: "Izin", value: stats.izin, color: "bg-red-500" },
        ].map((item, i) => (
          <div key={i} className={`${item.color} rounded-xl p-6 text-white shadow-lg transform hover:scale-105 transition-all`}>
            <p className="text-sm opacity-80">{item.label}</p>
            <p className="text-4xl font-bold mt-2">{item.value}</p>
          </div>
        ))}
      </div>
      <div className="bg-white rounded-xl p-6 shadow-lg">
        <h2 className="text-xl font-semibold mb-4">Aktivitas Terbaru</h2>
        <div className="space-y-3">
          {["Budi check-in pukul 08:00", "Ani check-out pukul 17:00", "Cici izin sakit", "Dodi terlambat 15 menit"].map((a, i) => (
            <div key={i} className="flex items-center gap-3 p-3 bg-gray-50 rounded-lg">
              <div className="w-2 h-2 bg-indigo-500 rounded-full"></div>
              <span className="text-gray-700">{a}</span>
            </div>
          ))}
        </div>
      </div>
    </div>
  );
}