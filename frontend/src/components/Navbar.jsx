import { Link, useLocation } from "react-router-dom";

const links = [
  { to: "/dashboard", label: "Dashboard", icon: "📊" },
  { to: "/absensi", label: "Absensi", icon: "✅" },
  { to: "/users", label: "Users", icon: "👥" },
];

export default function Navbar() {
  const location = useLocation();
  return (
    <nav className="bg-white shadow-lg sticky top-0 z-50">
      <div className="container mx-auto px-4">
        <div className="flex items-center justify-between h-16">
          <Link to="/" className="text-xl font-bold text-indigo-600">
            🏢 Sistem Absensi
          </Link>
          <div className="flex gap-2">
            {links.map((l) => (
              <Link
                key={l.to}
                to={l.to}
                className={`px-4 py-2 rounded-lg transition-all duration-200 ${
                  location.pathname === l.to
                    ? "bg-indigo-600 text-white shadow-md"
                    : "text-gray-600 hover:bg-indigo-50"
                }`}
              >
                {l.icon} {l.label}
              </Link>
            ))}
          </div>
        </div>
      </div>
    </nav>
  );
}