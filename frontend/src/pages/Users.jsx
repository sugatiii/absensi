import { useState } from "react";

const dummyUsers = [
  { id: 1, nama: "Budi Santoso", email: "budi@example.com", role: "Karyawan", status: "Aktif" },
  { id: 2, nama: "Ani Wijaya", email: "ani@example.com", role: "Admin", status: "Aktif" },
  { id: 3, nama: "Cici Putri", email: "cici@example.com", role: "Karyawan", status: "Izin" },
];

export default function Users() {
  const [search, setSearch] = useState("");
  const filtered = dummyUsers.filter(u => u.nama.toLowerCase().includes(search.toLowerCase()));

  return (
    <div className="space-y-6">
      <div className="flex items-center justify-between">
        <h1 className="text-3xl font-bold text-gray-800">User Management</h1>
        <button className="px-4 py-2 bg-indigo-600 text-white rounded-lg hover:bg-indigo-700 transition-all">+ Tambah User</button>
      </div>
      <div className="bg-white rounded-xl shadow-lg overflow-hidden">
        <div className="p-4 border-b">
          <input
            type="text"
            placeholder="🔍 Cari user..."
            value={search}
            onChange={(e) => setSearch(e.target.value)}
            className="w-full px-4 py-2 border rounded-lg focus:outline-none focus:ring-2 focus:ring-indigo-500"
          />
        </div>
        <table className="w-full">
          <thead className="bg-gray-50">
            <tr>
              {["Nama", "Email", "Role", "Status", "Aksi"].map(h => <th key={h} className="px-6 py-3 text-left text-sm font-semibold text-gray-600">{h}</th>)}
            </tr>
          </thead>
          <tbody className="divide-y">
            {filtered.map(u => (
              <tr key={u.id} className="hover:bg-gray-50 transition-all">
                <td className="px-6 py-4 font-medium">{u.nama}</td>
                <td className="px-6 py-4 text-gray-500">{u.email}</td>
                <td className="px-6 py-4">{u.role}</td>
                <td className="px-6 py-4"><span className={`px-2 py-1 rounded-full text-xs font-semibold ${u.status === "Aktif" ? "bg-green-100 text-green-700" : "bg-yellow-100 text-yellow-700"}`}>{u.status}</span></td>
                <td className="px-6 py-4"><button className="text-indigo-600 hover:underline">Edit</button></td>
              </tr>
            ))}
          </tbody>
        </table>
      </div>
    </div>
  );
}