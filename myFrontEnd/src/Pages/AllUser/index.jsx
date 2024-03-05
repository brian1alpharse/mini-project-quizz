import React, { useEffect, useState } from 'react';

const AllUser = () => {
  const [userRole, setUserRole] = useState('');

  useEffect(() => {
    // Mengambil token dari local storage
    const token = localStorage.getItem('token');

    // Lakukan pengecekan role user menggunakan token
    // Disesuaikan dengan logika autentikasi pada server Anda
    // Contoh sederhana: jika token ada, maka userRole = 'admin', jika tidak ada, maka userRole = 'user'
    if (token) {
      setUserRole('admin');
    } else {
      setUserRole('user');
    }
  }, []);

  return (
    <div>
      <h2>All User Page</h2>
      {userRole === 'admin' ? (
        <div>
          {/* Tampilkan data user untuk role admin */}
          <p>Data user untuk role admin</p>
        </div>
      ) : (
        <div>
          {/* Tampilkan pesan bahwa role user tidak memiliki akses */}
          <p>Maaf, Anda tidak memiliki akses sebagai user.</p>
        </div>
      )}
    </div>
  );
};

export default AllUser;
