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
            //var fm = new FileManager();
            //fm.Process("Files/01.txt");
            //Console.WriteLine(fm.LogText);

            var pm = new ParserManager();
            string fullString = "group = '2843360' and (team in ('TEAM1', 'TEAM2','TEAM3'))";
            pm.Process(fullString);
        }
    }
}
