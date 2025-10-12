import React from 'react';

const Home: React.FC = () => {
  return (
    <div className="min-h-screen bg-gradient-to-br from-blue-50 to-indigo-100 flex items-center justify-center">
      <div className="max-w-2xl mx-auto text-center p-8">
        <h1 className="text-5xl font-bold text-gray-900 mb-6">
          Welcome to Easy Ballot
        </h1>
        <p className="text-xl text-gray-600 mb-8">
          This is the home page of your ballot application.
        </p>
        <div className="space-x-4">
          <button className="bg-blue-600 hover:bg-blue-700 text-white font-semibold py-3 px-6 rounded-lg transition duration-200 shadow-lg">
            Get Started
          </button>
          <button className="bg-white hover:bg-gray-50 text-blue-600 font-semibold py-3 px-6 rounded-lg border-2 border-blue-600 transition duration-200">
            Learn More
          </button>
        </div>
      </div>
    </div>
  );
};

export default Home;
