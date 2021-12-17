using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace LoggingParserTest01
{
    class Program
    {
        static void Main()
        {
            var fm = new FileManager();
            fm.Process("Files/01.txt");
            Console.WriteLine(fm.LogText);
        }
    }
}
